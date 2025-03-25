package main

// You are given a 0-indexed binary string s having an even length.
//
// A string is beautiful if it's possible to partition it into one or more substrings such that:
//
// Each substring has an even length.
// Each substring contains only 1's or only 0's.
// You can change any character in s to 0 or 1.
//
// Return the minimum number of changes required to make the string s beautiful.

// Input: s = "1001"
// Output: 2
// Explanation: We change s[1] to 1 and s[3] to 0 to get string "1100".
// It can be seen that the string "1100" is beautiful because we can partition it into "11|00".
// It can be proven that 2 is the minimum number of changes needed to make the string beautiful.

// how to prove that?
// Input: 1001
// I can only make two partition because a partition needs to be even in length. First partition is 2 / 2 in length
// So only possible partition is 10 / 01
// each partition needs to be made beautiful
// 10 => either 11 or 00 (1 change)
// 01 => either 11 or 00 (1 change)
// second partition is 0 / 4 = 4 / 0 and it also needs 2 changes to be beautiful: 1001 => 1111 or 0000.
// so both partitions need 2 changes, hence the minimum is 2.

// hypothesis: if a partition needs N changes to be beautiful, splitting it to X subpartitions cannot be all beautiful in M < N changes
// for instance: 000001 needs one change
// if I split it into three partitions: 00 00 01
// I still need one change

// in other words, if a 2-length partition is not beautiful, there needs to be at least one change to make that partition beautiful

// other sample: 001000
// 001 / 000 => 1 change
// 00 / 10 / 00 => 1 change

// other sample: 001010
// 001 / 010 => 2 changes
// 00 10 10 => 2 changes as well

func minChanges(s string) int {
	count := 0
	for i := 0; i < len(s)/2; i++ {
		j := i * 2
		if s[j] != s[j+1] {
			count++
		}
	}
	return count
}
