* In some implementations, array of boolean is used instead of array of bytes, which is not an optimized solution.
* A boolean takes 1 byte ==> 8 bits in Golang
* int => 32 bits => 4 bytes
* First try to clear the input and output in your mind.
    * Bloom filter will have an Add method to check if something exists and Exists method to check if that element is present.
* Let's start with simpler implementation of array of boolean and then we can optimize it.
* Hash algorithm to be used for bloom filters:
    * Speed
    * Distribution: Uniform distribution to minimize collision.
    * Avalanche effect: A small change in input should result in a significant change in the hash value. This helps in reducing false positive. how?
        * Even distribution
        * Diffusion of information: Bits affected by the original hash will be completely different from the bits affected by the modified hash. Reduces the number of shared bits.
        * Unpredictability: increases the security and reliability.
    * Non-cryptographic: Unnecessarily reduces speed.
* Other factors affecting performance and effectiveness of the bloom filter:
    * Number of hash functions used.
    * Size of the byte array.
    * Number of elements being stored in the filter.
* murmur3 hash github golang
* hash.Hash has a New32WithSeed, Write, Sum32, Reset function.
    * Why a Reset? Else it will not remove the written data and instead data will be appended on the existing data.

* *False positive*: Our system said yes but it was a no.
* *False negative*: Our system said no but it was actually present.
    * This is zero in bloom filters.
* *How false positivity increases when we keep adding more data to bloom filter*
    * Let's create some dataset to see that.
    * Inflection point: concave up to down or vice-versa.
    * How we divide the data in data science?
        * Train and test, we'll do something similar.
* *What should be the size of the bloom filter? After some point, increasing the size of the bloom filter won't help much, we need to find that point.*

* *False positivity rate*:
    * % of cases where 
    * Have a dataset called datasetNotExists ==> none of them should be present in the bloom filter. ==> this is FP + TN
    * We have 2 datasets of the same size.
    * We add elements of the first dataset to the bloom filter.
    * We iterate through the elements of the second dataset. If any of them turns to be present in the bloom filter, then it is a false positive.
    * You have a dataset.
    * You added elements into it.
    * There won't be any FN in bloom filter.
    * Hence dataset size = FP + TN

* *Using Bit Manipulation*
    * int8 ==> set 17th bit ==> 3rd number, 2 element
        ==> array element ==> i / 8, bit position ==> i % 8.


* *FPR: Increasing number of hash functions*
    * This will make it more accurate at the start because we will consider that a key exists only if all of the hash functions has their respective idx set.
    * After some point, from decreasing fpr, this will start increasing fpr (we need to find the sweet spot).
        * This is because, with increasing hash functions, the chances of collision also increases.