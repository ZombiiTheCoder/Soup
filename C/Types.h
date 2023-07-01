#ifndef TYPES_H
#define TYPES_H

typedef FILE                File;
typedef size_t              st;
typedef ptrdiff_t           dt;

typedef long int            lint;
typedef unsigned char       uchar;
typedef unsigned int        uint;
typedef unsigned long       ul;
typedef unsigned long long  ull;

typedef uint8_t             u8;
typedef uint16_t            u16;
typedef uint32_t            u32;
typedef uint64_t            u64;

typedef int8_t              i8;
typedef int16_t             i16;
typedef int32_t             i32;
typedef int64_t             i64;

#define I8MIN               INT8_MIN
#define I8MAX               INT8_MAX
#define I16MIN              INT16_MIN
#define I16MAX              INT16_MAX
#define I32MIN              INT32_MIN
#define I32MAX              INT32_MAX
#define I64MIN              INT64_MIN
#define I64MAX              INT64_MAX

#define U8MAX               UINT8_MAX
#define U16MAX              UINT16_MAX
#define U32MAX              UINT32_MAX
#define U64MAX              UINT64_MAX
#define PI                  M_PI

#endif /* end of include guard: TYPES_H */