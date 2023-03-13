# Lossless-JPEG
A revisit to the Lossless JPEG codec and rewriting it in golang. 

## What is Lossless-JPEG
This is a not very well known (to the general community) image compression algorithm that does not share much similarities with the popular JPEG besides 
being originated from it. It is a predictive algorithm that compresses images in a way that is 100% lossless.

This codec came to be somewhat popular for medical imaging and in some cameras for compressing raw images.

**Note:** Just so you know, JPEG with 100 quality is, surprisingly, not actually lossless. It loses a tiny bit of data each time it is compressed. Up to
5%, to be specific. 

**Beware:** There are other algorithms out there that are able to compress images in a 100% lossless way and this may not be the most optimised solution
in terms of how much the image can be compressed.

## Oh so this is what JPEG-LS stands for, right?
> ***NO.***

JPEG-LS and Lossless JPEG are 2 different algorithms, contrary to what some unreliable sources will tell you out there. Lossless JPEG was created
on top of JPEG in 1993.

# Functioning of this Codec
> TLDR: **Original image** -> **DPCM Predictor** -> **JPEG encoding** -> **Huffman encoding** -> **Compressed image**

I will update this with the details of the process at a later time.
