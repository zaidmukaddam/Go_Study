# hash-cracker-GoLang

Simple multithreading brute hash **MD5**, **SHA128**, **SHA256** and **SHA1**.

## Usage
    ./dehash -hash hexstring|-file path [-type sha1|sha256|sha512|md5] [-max val] [-min val] [-charset chars] [-threads num] 

    Usage of ./dehash:
    -charset string
          char set for possible message (default "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\v\f")
    -file string
          path to file with hashes
    -hash string
          hash to be decrypted
    -max int
          max length of message (default 10)
    -min int
          min length of message (default 1)
    -threads int
          max number of threads (default 256)
    -type string
          hash algorithm (default "sha1")


## Setup
Clone the repository and change the working directory:

    git clone https://github.com/zaidmukaddam/Go_Study/tree/main/day14/hash-cracker-GoLang
    cd hash-cracker-GoLang

Build and run the program:
    
    go build -o dehash
    ./dehash -hash aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d -charset abcdefghijklmnopqrstuvwxyz

## Examples

Crack **sha1** hash
  
    ./dehash -hash fc19318dd13128ce14344d066510a982269c241b

Crack **md5** hash with message length between 3 and 5 characters
    
    ./dehash -type md5 -min 2 -max 5 -hash 6b2ded51d81a4403d8a4bd25fa1e57ee

Crack **sha1** hash with custom char set
    
    ./dehash -hash b3da1b8c56c939e616aa3c0934bce72cb1e82b32 -charset abcdhijklm

Crack **sha256** with 512 threads
    
    ./dehash -type sha256 -threads 512 -hash 97c10efe01d5c9c88704a12d361d8429b3a6aa2412290a0773109d5d2d603d5e

Crack **sha1** hashes from file
    
    ./dehash -type sha1 -file ./hashes.txt


## Output examples

Ex. 1

    $ ./dehash -hash fc19318dd13128ce14344d066510a982269c241b
    Start cracking hash fc19318dd13128ce14344d066510a982269c241b
    Check messages with length: 1 | Possible variants: 100
    Check messages with length: 2 | Possible variants: 10000
    Check messages with length: 3 | Possible variants: 1000000
    Check messages with length: 4 | Possible variants: 100000000
    =========> Message: good

Ex. 2 

    $ ./dehash -type sha256 -min 5 -max 9 -file hashes.txt -charset abcdehijklmnt
    Start cracking hash 1c8bfe8f801d79745c4631d09fff36c82aa37fc4cce4fc946683d7b336b63032
    Check mesages with length: 5 | Possible variants: 371293
    Check mesages with length: 6 | Possible variants: 4826809
    Check mesages with length: 7 | Possible variants: 62748517
    =========> Message: letmein

Ex. 3

    $ ./dehash -max 5 -file ./hashes.txt -charset abcdefghijklmnopqrstuvwxyz/:.
    Start cracking hash c3437dbc7c1255d3a21d444d86ebf2e9234c22bd
    Check messages with length: 1 | Possible variants: 29
    Check messages with length: 2 | Possible variants: 841
    Check messages with length: 3 | Possible variants: 24389
    Check messages with length: 4 | Possible variants: 707281
    Check messages with length: 5 | Possible variants: 20511149
    =========> Message: https

    Start cracking hash ef81042e1e86acb765718ea37393a1292452bbcc
    Check messages with length: 1 | Possible variants: 29
    Check messages with length: 2 | Possible variants: 841
    Check messages with length: 3 | Possible variants: 24389
    =========> Message: ://

    Start cracking hash a3c1509bd8df6d72992b312e4f6b7f4ce7fd3f3d
    Check messages with length: 1 | Possible variants: 29
    Check messages with length: 2 | Possible variants: 841
    Check messages with length: 3 | Possible variants: 24389
    Check messages with length: 4 | Possible variants: 707281
    Check messages with length: 5 | Possible variants: 20511149
    =========> Message: docs.

    Start cracking hash 3f95edc0399d06d4b84e7811dd79272c69c8ed3a
    Check messages with length: 1 | Possible variants: 29
    Check messages with length: 2 | Possible variants: 841
    Check messages with length: 3 | Possible variants: 24389
    =========> Message: goo
