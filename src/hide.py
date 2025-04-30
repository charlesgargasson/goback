import sys
mystring=sys.argv[1]
charsets = {
    'x':["o","y","a","e","i","u"],
    'X':["U","I","Y","O","A","E","/","\\",':']
}
delim = '`'

lastcharwasincharset = None
def printChar(char, charset, ischarset):
    global lastcharwasincharset
    if ischarset :
        if lastcharwasincharset == False:
            sys.stdout.write(delim)
        if lastcharwasincharset != None:
            sys.stdout.write(r'+')
        sys.stdout.write(charset+'[')
        sys.stdout.write(str(charsets[charset].index(char)))
        sys.stdout.write(r']')
        lastcharwasincharset = True
    else:
        if lastcharwasincharset == None:
            sys.stdout.write(delim)
        if lastcharwasincharset:
            sys.stdout.write(r'+')
            sys.stdout.write(delim)
        sys.stdout.write(char)
        lastcharwasincharset = False

for char in mystring:
    found = False
    for charset, charsetlist in charsets.items():
        if char in charsetlist:
            printChar(char,charset,True)
            found = True
            continue

    if not found :
        printChar(char,None,False)

if not lastcharwasincharset:
    sys.stdout.write(delim)
sys.stdout.write('\n')
