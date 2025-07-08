
func categorizeBox(length int, width int, height int, mass int) string {
    var bulky,heavy,answer string
    if (length >= 10000 || width >= 10000 || height >= 10000) ||  (length*width*height) >= 1000000000  {
        bulky="Bulky"
    }
    if mass >= 100{
        heavy="Heavy"
    }

    switch {
        case bulky == "Bulky" && heavy == "Heavy" : {
          answer="Both"

        }
                case bulky == "Bulky" && heavy == "Heavy" : {
          answer="Both"

        }
        case bulky != "Bulky" && heavy != "Heavy" : {
          answer="Neither"

        }
                case bulky == "Bulky" && heavy != "Heavy" : {
          answer="Bulky"

        }
                case bulky != "Bulky" && heavy == "Heavy" :{
          answer="Heavy"

        }



    }
    return answer
    
}