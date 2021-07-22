def convert(number):
    rain = ""
    rain += "Pling" if number % 3 == 0 else ""
    rain += "Plang" if number % 5 == 0 else ""
    rain += "Plong" if number % 7 == 0 else ""
    #* Check if the string is empty
    if not rain:
        return str(number)
    return rain