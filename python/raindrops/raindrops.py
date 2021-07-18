def convert(number):
    rain = ""
    factor_three = number % 3 == 0
    factor_five = number % 5 == 0
    factor_seven = number % 7 == 0
     
    if factor_three:
        rain += "Pling"
    if factor_five:
        rain += "Plang"
    if factor_seven:
        rain += "Plong"
    #* Check if the string is empty
    if not rain:
        return str(number)
    return rain