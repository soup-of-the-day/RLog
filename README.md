# Rudiment Practice Log (RLog)
I am a drummer. I am also a lazy drummer, with a bad habit. I devote too much practice time to skills that don't make me 
that much better of a drummer. Too many dumb fills, worrying too much about speed. I need to focus on fundamental sticking 
skills to improve (I've hit a plateau, oh no). 

So, this tool is designed to track my progress on my rudiments while also selecting random rudiments for me to practice.

## Usage

```
./RLog start <number of rudiments to practice>
```
This will `start` a practice session and return to you how ever many rudiments you wish to practice. These are randomly 
chosen from the `rudiments.csv` file present in the directory that RLog is called from.

```
./RLog end [ (<3-letter abbreviation of rudiment>, <BPM> ) ... ]
``` 
This will `end` a practice session, and update the rudiments in `rudiments.csv` with their new speeds. You can pass in 
any number of rudiments to update, example:

`./RLog end ssr 78 dsr 90` -> Will update your single stroke roll to be 78 BPM and your double stroke roll to be 90 BPM

You can find the 3-letter acronyms for each stroke in `rudiments.csv`, but here is a list for your convenience:

```
Single Stroke Roll -> ssr
Single Stroke Four -> ss4
Single Stroke Seven -> ss7
Three Stroke Roll -> 3sr
Five Stroke Roll -> 5sr
Six Stroke Roll -> 6sr
Seven Stroke Roll -> 7sr
Nine Stroke Roll -> 9sr
Ten Stroke Roll -> 10sr
Eleven Stroke Roll -> 11sr
Thirteen Stroke Roll -> 13sr
Fifteen Stroke Roll -> 15sr
Seventeen Stroke Roll -> 17sr
Single Paradiddle -> spd
Double Paradidle -> dpd
Triple Paradiddle -> tpd
Single Paradiddle-diddle -> spdd
```