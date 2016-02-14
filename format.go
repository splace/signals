package signals

import (
	"encoding/binary"
	"fmt"
	"io"
)

// encode as unsigned PCM data in a Riff wave container (wav file format) 
func Encode(w io.Writer, s Signal, length interval, sampleRate uint32, bytes uint8) {
	binaryWrite := func(w io.Writer, d interface{}) {
		if err := binary.Write(w, binary.LittleEndian, d); err != nil {
			panic(err)
		}
	}
	samplePeriod := MultiplyInterval(1/float32(sampleRate), UnitTime)
	samples := uint32(length/samplePeriod) + 1
	fmt.Fprint(w, "RIFF")
	binaryWrite(w, uint32(samples+36))
	fmt.Fprint(w, "WAVE")
	fmt.Fprint(w, "fmt ")
	binaryWrite(w, uint32(16))
	binaryWrite(w, uint16(1))
	binaryWrite(w, uint16(1))
	binaryWrite(w, sampleRate)
	binaryWrite(w, sampleRate)
	binaryWrite(w, uint16(bytes))
	binaryWrite(w, uint16(8*bytes))
	fmt.Fprint(w, "data")
	binaryWrite(w, uint32(samples))
	var i uint32
	var shift uint8 = LevelBits - 8*bytes
	var offset level = 2 << (bytes*8 - 2)
	for ; i < samples; i++ {
		w.Write([]byte{byte(s.Level(interval(i)*samplePeriod)>>shift + offset)})
	}
}

/* PCM all possible formats

openal takes s16 or u8

 DE alaw            PCM A-law
 DE f32be           PCM 32-bit floating-point big-endian
 DE f32le           PCM 32-bit floating-point little-endian
 DE f64be           PCM 64-bit floating-point big-endian
 DE f64le           PCM 64-bit floating-point little-endian
 DE mulaw           PCM mu-law
 DE s16be           PCM signed 16-bit big-endian
 DE s16le           PCM signed 16-bit little-endian
 DE s24be           PCM signed 24-bit big-endian
 DE s24le           PCM signed 24-bit little-endian
 DE s32be           PCM signed 32-bit big-endian
 DE s32le           PCM signed 32-bit little-endian
 DE s8              PCM signed 8-bit
 DE u16be           PCM unsigned 16-bit big-endian
 DE u16le           PCM unsigned 16-bit little-endian
 DE u24be           PCM unsigned 24-bit big-endian
 DE u24le           PCM unsigned 24-bit little-endian
 DE u32be           PCM unsigned 32-bit big-endian
 DE u32le           PCM unsigned 32-bit little-endian
 DE u8              PCM unsigned 8-bit

*/

