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

    LRT_Table *LRT_NewTable();            // Create a new Table.
    void LRT_DropTable(LRT_Table *table); // Drop the Table.

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
    // Adds all elements of one Table to another.
    void LRT_TableAddAll(LRT_Table *to, LRT_Table *from);

#ifdef __cplusplus
}
#endif

#endif