package hash

// NOTE: you can ignore this file

// IsHash returns true if the index is a properly created hash index
func IsHash(index *HashIndex) (bool, error) {
	table := index.GetTable()
	buckets := table.GetBuckets()
	for _, pn := range buckets {
		// Get bucket
		bucket, err := table.GetBucket(pn)
		d := bucket.GetDepth()
		if err != nil {
			return false, err
		}
		// Get all entries
		entries, err := bucket.Select()
		if err != nil {
			return false, err
		}
		// Check that all entries should hash to this bucket.
		for _, e := range entries {
			key := e.Key
			hash := Hasher(key, d)
			if pn != table.buckets[hash] {
				return false, nil
			}
		}
	}
	return true, nil
}
