package generator





    
type S struct{
    
        E *E
    
}


    
type E struct{
    
        T *T
    
        ; *Terminal
    
        + *Terminal
    
        E *E
    
}


    
type T struct{
    
        E *E
    
        ) *Terminal
    
        Int *Terminal
    
        ( *Terminal
    
}


    
type Terminal struct{
    
        Value string
    
}





    
func  ContructParserNode(){
    
    




if root.String()==S {
    
    




if right==E {
    
} 
     





    
}  else 
if root.String()==E {
    
    




if right==T {
    
} 
     





    
    




if right==; {
    
} 
     





    
    




if right==T {
    
} 
     





    
    




if right==+ {
    
} 
     





    
    




if right==E {
    
} 
     





    
}  else 
if root.String()==T {
    
    




if right==( {
    
} 
     





    
    




if right==E {
    
} 
     





    
    




if right==) {
    
} 
     





    
    




if right==Int {
    
} 
     





    
} 
 
 
     





    
}























