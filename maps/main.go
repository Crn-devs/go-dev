package main

/**
	Maps in Go (Key/Value)

	a map is go's built in hashtable

	reference type

	internally in go maps headers contain
		pointer to hash table data (buckets and overflow buckets)
		count: the number of key value pairs
		other runtime bookkeeping: hashseed, bucket count, etc

	internally maps are structs that hold pointers to an hmap in the runtimes memory, this isnt super important
	for general use but nice to know, hmap cannot be imported and is not part of the language spec
	so from the language point of view a map is a abstract reference type with no fields you can access

	extra info:
		at runtime a map is an hmap in the package runtime/map.go pointing to an array of buckets plus overflow buckets if any

		they work quite interestingly im going to write out the high level steps of what happens when we
		create a map, add a key, and lookup a key

		high level overview of map creation process in go

			m := make(map[string]int, 100)
			under the hood this line of code does the following steps
				allocates an internal hmap (this is the control struct)
				seeds it with a random hash seed
				decides an inital bucket count 2^B from the size hint
					if we create a map with make(..., hint) -> the runtime choses B so 2^B = hint/load factor
					with no hint or a empty literal then buckets may allocated lazily on the first insert

			buckets - fixed size structures that store the tophash, the actual key and their values, an overflow pointer for extra slots

			conceptually a bucket looks like this
					tophash: [ t0 t1 t2 t3 t4 t5 t6 t7 ]   // 1 byte per slot (hi 8 bits of hash, offset by +4)
					keys:    [ k0 k1 k2 k3 k4 k5 k6 k7 ]   // 8 keys packed contiguously
					values:  [ v0 v1 v2 v3 v4 v5 v6 v7 ]   // 8 values packed contiguously
					overflow: *bmap                         // pointer to another bucket if full (or nil)

			Tophash stores only a partial hash, high 8bits +4 not the full hash
			keys and values are laid out contiguously right after the tophash bites which is great for cache locality
			if all 8 slots are used the bucket links to an overflow bucket


			here is a high level overview of what a insert may look like

			1. hash
			2. Choose a bucket (low bit mask)
			3. compute the tophash
			4. scan bucket
				1. walk the 8 top hash bytes
					if tophash[i] == 0 remember the first empty slot -> not found yet -> keep a reference of the slot for insertion if not found
					if tophash[i] == top -> compare the key at slot i
						if equal overwrite the value and done
				2. place or overflow
					if an empty slot is found write the key/value and tophash[i] = top
					if no empty slot is found then follow/create the overflow and re-scan there
				3. Potentially trigger growth

		what makes hashmaps (maps) fast

			O(1)avg via hash -> mask -> small array scan
			tophash byte filters most non matches before having to check long key/value
			cache friendly layout

		pitfalls in go maps
			not concurrency safe, must use sync.RWMutex or sync.Map
			deleting doesnt shrink the map it uses tombstones and doesnt reduce the allocated buckets
				rebuild the map if we need to de-allocate/compact
			unordered: maps dont Keep insertion order if we need stable order collect the keys and sort them


			what is a map in go ?

			“A Go map is an hmap pointing to a power-of-two array of buckets.
			 Each bucket has 8 slots: a tophash byte per slot, then packed keys and values,
			  and maybe an overflow link. To insert or lookup, we hash the key with the map’s seed,
			  mask the low B bits to pick a bucket, then scan that bucket’s 8 tophash bytes
			   to quickly find candidates. Only on tophash matches do we compare the real key.
			    If a bucket is full, we follow or allocate an overflow bucket. When the table is too full,
				 Go doubles the number of buckets and incrementally evacuates
				 old buckets so we never pause for a full rehash.”

**/

func main() {

}
