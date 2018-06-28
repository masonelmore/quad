# quad
Generates interesting pictures by using a bastardized implementation of a quadtree to group colors into squares based on how far away their average color is from their region.

# usage
1. Install Go https://golang.org/dl/
2. Install quad `go get github.com/masonelmore/quad`
3. Run quad `# $GOROOT\bin\quad <filename>.png tolerance`

### notes and limitations
* PNG is currently the only supported filetype.
* Tolerance is a float between 0 and 1.  Lower values produce more detailed images.


# examples
|original|quadified|tolerance|
|-|-|-|
|![obama original](images/obama.png)|![obama quad](images/obama-0.045.png)|0.045|
|![park original](images/park.png)|![park quad](images/park-0.06.png)|0.06|
|![sunset original](images/sunset.png)|![sunset quad](images/sunset-0.06.png)|0.06|
