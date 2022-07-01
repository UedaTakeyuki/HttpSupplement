package HttpSupplement

/** 
 Convert Float64 type data to Int64
   
   An Int64 data get converted to Float64 by Json marshaller.
   So, convert back to Int64.
*/
func ConvertToInt64IfFloat64(in interface{})(out interface{){
  	switch v := in.(type) {
		case float64:
			out = int64(v)
		default:
			out = in
		}
    return
}
