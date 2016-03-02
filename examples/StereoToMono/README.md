```
Usage:
tomono.(exe|bin) <<flags>> <<inFile.wav>> <<outfile>>
 -bytes precision
    	precision in bytes per sample. (requires format option set) (default 2)
  -chans string
    	extract/recombine listed channel number(s) only. ('1,2' for first 2 channels) (default "1,2")
  -db uint
    	adjust volume in dB (-6 to halve.) stacked channels could clip without.
  -format
    	don't use input sample rate and precision for output, use command-line options
  -help
    	display help/usage.
  -prefix string
    	add individual prefixes to extracted mono file(s) names. (default "L-,R-,C-,LFE-,LB-,RB-")
  -rate samples
    	samples per second.(requires format option set) (default 44100)
  -stack
    	recombine all channels into a mono file.

```
