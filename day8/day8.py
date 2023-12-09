

def gcd(a, b):
    while b:
        a, b = b, a % b
    return a


def gcd_array(numbers):
    result = numbers[0]
    for num in numbers[1:]:
        result = gcd(result, num)
    return result


def lcm(a, b):
    return abs(a * b) // gcd(a, b) if a and b else 0


def lcm_array(numbers):
    result = 1
    for num in numbers:
        result = lcm(result, num)
    return result


# Example usage:
numbers = [13019, 14681, 20221, 19667, 18559, 16897]
numbers2 = [40442, 50691, 39057, 44043, 39334, 55677]
result = lcm_array(numbers2)
print(f"The GCD of {numbers} is: {result}")
