#ifndef LOXCRT_TABLE_H
#define LOXCRT_TABLE_H

#include "prelude.h"
#include "object.h" // for StringObject
#include "value.h"  // for Value

#ifdef __cplusplus
extern "C"
{
#endif

/**
 * The load factor of (hash) Table.
 *
 * Load factor is the number of entries divided by the number of buckets. The lower the
 * factor is, the less the chance of collisions.
 */
#define TABLE_MAXLOAD 0.75

    /**
     * The entry of a Table.
     *
     * The table is a similar data structure to `dict[str, Value]` in Python, or
     * `HashMap<String, Value>` in Java. Since C does not support generics, we have to
     * decide the type of keys and values where defined.
     *
     * Luckily, LOXCRT does not use a lot of Tables. String interning uses one, and each
     * instance object is actually a table.
     */
    typedef struct
    {
        LRT_StringObject *key;
        LRT_Value value;
    } LRT_TableEntry;

    // The hashmap with keys typed StringObject, and values typed Value.
    //
    // This is used in string interning, and as instance objects.
    typedef struct
    {
        size_t length;
        size_t capacity;
        LRT_TableEntry *entries;
    } LRT_Table;

    void LRT_InitializeTable(LRT_Table *table);
    void LRT_FinalizeTable(LRT_Table *table);

    /**
     * Lookup the corresponding value of a specific key.
     *
     * If there exists the given key, the value is passed by the pointer parameter `value`
     * and a `true` is returned; otherwise, nothing is written to `value` and a `false`
     * value is returned.
     */
    bool LRT_TableGet(LRT_Table *table, LRT_StringObject *key, LRT_Value *value);
    /**
     * Insert a key-value pair into the Table.
     *
     * If there's already a same key, the old value is overwritten; otherwise, a new entry
     * is added and a `true` value is returned.
     */
    bool LRT_TableSet(LRT_Table *table, LRT_StringObject *key, LRT_Value value);
    /**
     * Delete an entry from the Table.
     *
     * Actually, we don't set the entry to the NIL value; that may misguide the TableGet
     * to stop searching neighboring entries if hash conflicts happen. Instead, we'll set
     * the key to NULL and the value to BOOLEAN(true), which indicates this is a tombstone
     * value.
     *
     * If the corresponding entry exists, a `true` is returned.
     */
    bool LRT_TableDelete(LRT_Table *table, LRT_StringObject *key);
    /**
     * Check if the Table contains the specified string as a key.
     *
     * If so, the pointer to the object is returned; otherwise NULL. This is vital for
     * string interning.
     */
    LRT_StringObject *LRT_TableContainsKey(LRT_Table *table, const char *chars, size_t length, uint32_t hash);
    // Adds all elements of one Table to another.
    void LRT_TableAddAll(LRT_Table *to, LRT_Table *from);

#ifdef __cplusplus
}
#endif

#endif