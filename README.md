# ASCII-WEB-STLIZE
This project is written in [Golang](https://go.dev/)
## Pre-requisite
go version 1.23.0 or higher



This project is a build up of **[ASCII-WEB](https://learn.zone01kisumu.ke/git/lakoth/ascii-art-web)**

Ascii-art-stylize consists on making your site :

    1.more appealing, interactive and intuitive.
    
    2.more user friendly.

    3.give more feedback to the user.

To clone the project 
```bash
git clone  https://learn.zone01kisumu.ke/git/lakoth/ascii-art-web-stylize
```
cd into the project directory with the command
```bash
cd ascii-art-stylize
```
Run the project
```bash
go run .
```
Click  the link displayed on the terminal (localhost:5000) and interract with the web-based version.

## Implementation Details

**ASCII Art Generation Algorithm**

The ASCII art generation process follows these steps:

* Banner Format:
        Each character is represented by an 8-line pattern in the banner style.
        Characters are separated by new lines (\n).

* User Input Processing:
        The application takes user input, which includes the text and the chosen banner style.

* Loading Banner Styles:
        The chosen banner style's ASCII art template is loaded from the corresponding text file in the "banners" directory.
        copy the files to your working directory
        **Download banner files**<br>
                [shadow.txt](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)<br>
                [thinkertoy.txt](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/thinkertoy.txt)<br>
                [standard.txt](https://learn.zone01kisumu.ke/git/root/public/src/branch/master/subjects/ascii-art/standard.txt)<br>

* Text to ASCII Art Conversion:
        The input text is split into lines, where each line represents a line of ASCII art.
        For each line of text:
            The line is split into individual characters.
            For each character:
                The character's ASCII value is checked to ensure it falls within the Basic Latin character range (ASCII values 32 to 126).
                If the character is within this range, its corresponding pattern from the banner style's ASCII art template is retrieved.
                The patterns of all characters in the line are concatenated to form the ASCII art representation for that line.
                The process is repeated for each of the 8 lines in the character's pattern.
                The application maps each character to its corresponding pattern in the banner style, resulting in the ASCII art representation for each line of text.

* Rendering the Result:
        The generated ASCII art, along with the original input text and chosen banner style, is rendered on the result page for the user to view.




## Contribute
To contribute make changes and submit a pull request.

## Contributers

**[Lorna Akoth](https://github.com/akothouma)**

**[Franklyne Namayi](https://github.com/fnamayi)**

