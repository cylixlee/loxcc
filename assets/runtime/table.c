#include "table.h"
#include <string.h> // for string manipulation
#include "gc.h"     // for memory allocation

/**
 * Algorithm for calculating new size of extended containers.
 *
 * Most dynamic array implementations have a minimum threshold like this. According to
 * amortized analysis, as long as we grow the array by a multiple of its current size,
 * when we average out the cost of a sequence of appends, each append is O(1).
 */
static inline size_t LRT_GrowCapacity(size_t capacity)
{
    if (capacity < 8)
    {
        return 8;
    }
    return capacity * 2;
}

/**
 * Find the entry of a specific key, no matter empty or not.
 *
 * Note that this function does not receive the whole LRT_Table struct as parameter, but
 * two separate `entries` and `capacity`. This is because, when the Table expands, all of
 * its elements are re-placed (because the capacity matters when indexing from the hash
 * value). Passing the two arguments provides more flexibility and reusability.
 */
static LRT_TableEntry *LRT_FindEntry(LRT_TableEntry *entries, size_t capacity, LRT_StringObject *key);
// Extend the table to reserve space for new elements.
static void LRT_AdjustCapacity(LRT_Table *table, size_t newCapacity);

void LRT_InitializeTable(LRT_Table *table)
{
    table->length = 0;
    table->capacity = 0;
    table->entries = NULL;
}

void LRT_FinalizeTable(LRT_Table *table)
{
    FREE(table->entries, LRT_TableEntry, table->capacity);
    LRT_InitializeTable(table); // zero-initialize
}

bool LRT_TableGet(LRT_Table *table, LRT_StringObject *key, LRT_Value *value)
{
    if (table->length == 0)
    {
        return false;
    }
    LRT_TableEntry *entry = LRT_FindEntry(table->entries, table->capacity, key);
    if (entry->key == NULL)
    {
        return false;
    }
    *value = entry->value;
    return true;
}

bool LRT_TableSet(LRT_Table *table, LRT_StringObject *key, LRT_Value value)
{
    // reserve enough space to meet the load factor
    if (table->length + 1 > table->capacity * TABLE_MAXLOAD)
    {
        size_t newCapacity = LRT_GrowCapacity(table->capacity);
        LRT_AdjustCapacity(table, newCapacity);
    }

    // find an entry and judge if it's a new key
    LRT_TableEntry *entry = LRT_FindEntry(table->entries, table->capacity, key);
    bool isNewKey = entry->key == NULL;
    if (isNewKey && IS_NIL(entry->value))
    {
        table->length++;
    }

    // set the key-value pair and returns
    entry->key = key;
    entry->value = value;
    return isNewKey;
}

bool LRT_TableDelete(LRT_Table *table, LRT_StringObject *key)
{
    if (table->length == 0)
    {
        return false;
    }
    LRT_TableEntry *entry = LRT_FindEntry(table->entries, table->capacity, key);
    if (entry->key == NULL)
    {
        return false;
    }
    entry->key = NULL;
    entry->value = BOOLEAN(true);
    return true;
}

LRT_StringObject *LRT_TableContainsKey(LRT_Table *table, const char *chars, size_t length, uint32_t hash)
{
    if (table->length == 0)
    {
        return NULL;
    }

    size_t index = hash % table->capacity;
    for (;;)
    {
        LRT_TableEntry *entry = &table->entries[index];
        if (entry->key == NULL)
        {
            // we don't immediately return NULL because this may be a tombstone.
            if (IS_NIL(entry->value))
            {
                // if it's indeed an empty slot, returns NULL because we haven't found the
                // key.
                return NULL;
            }
        }
        else if (entry->key->length == length && entry->key->hash == hash)
        {
            // otherwise, we compare the length and hash first, because number comparison
            // is much faster than strcmp.
            if (strcmp(entry->key->chars, chars) == 0)
            {
                // and if the strings are equal too, we've now found one and return it.
                return entry->key;
            }
        }
        index = (index + 1) % table->capacity;
    }
}

void LRT_TableAddAll(LRT_Table *to, LRT_Table *from)
{
    for (size_t i = 0; i < from->capacity; i++)
    {
        LRT_TableEntry *entry = &from->entries[i];
        if (entry->key == NULL)
        {
            continue;
        }
        LRT_TableSet(to, entry->key, entry->value);
    }
}

static LRT_TableEntry *LRT_FindEntry(LRT_TableEntry *entries, size_t capacity, LRT_StringObject *key)
{
    size_t index = key->hash % capacity; // find index according to the hash
    LRT_TableEntry *tombstone = NULL;
    for (;;)
    {
        LRT_TableEntry *entry = &entries[index];
        if (entry->key == NULL)
        {
            if (IS_NIL(entry->value))
            {
                // if we've found a tombstone which matches the hash, just use it;
                // otherwise, we return the empty entry for Get or Set.
                return tombstone != NULL ? tombstone : entry;
            }
            else
            {
                // if we've found a tombstone for the first time, record it.
                if (tombstone == NULL)
                {
                    tombstone = entry;
                }
            }
        }
        // linear probe until an empty entry is found.
        index = (index + 1) % capacity;
    }
    // the loop will definitely end because we've adjusted the capacity before finding an
    // entry.
}

static void LRT_AdjustCapacity(LRT_Table *table, size_t newCapacity)
{
    // allocate and zero-initialize entries
    LRT_TableEntry *newEntries = ALLOCATE(LRT_TableEntry, newCapacity);
    for (size_t i = 0; i < newCapacity; i++)
    {
        newEntries[i].key = NULL;
        newEntries[i].value = NIL;
    }

    // re-place existing elements
    //
    // Note that we cannot write `i < table->length` because the elements in the table is
    // (probably) not continuously stored. Moreover, all existing entries are re-indexed
    // because capacity matters in indexing entries (from the hash value), and now the
    // capacity has changed.
    //
    // Tombstones are cleared when adjusting capacity. Reset the length.
    table->length = 0;
    for (size_t i = 0; i < table->capacity; i++)
    {
        LRT_TableEntry *source = &table->entries[i];
        if (source->key == NULL)
        {
            continue;
        }
        LRT_TableEntry *destination = LRT_FindEntry(newEntries, newCapacity, source->key);
        destination->key = source->key;
        destination->value = source->value;
        table->length++;
    }

    // free the old entries and fill it with the new one.
    FREE(table->entries, LRT_TableEntry, table->capacity);
    table->entries = newEntries;
    table->capacity = newCapacity;
}