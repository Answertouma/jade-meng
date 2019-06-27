namespace go love

struct LoveStruct {
    i32 Id
    string Woqu
}

typedef i32 ouhou

service Woqu {
    i32 getId()
    LoveStruct getStruct()
}

service Woqule {
    i32 getIds()
    string getName()
}