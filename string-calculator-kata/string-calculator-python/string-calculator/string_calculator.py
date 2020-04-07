
def string_calculator(numbers):
    return 0

# Depending on the tests, you would probably inherit BaseGAETest, VTestCase or something similar here
# Those classes include methods you can override to help with testing (ex: setUp methods that run before your tests)
class StringCalculatorTests():
    def should_return_0_on_empty_string(self):
        val = string_calculator('')
        assert(val, 0)

    def should_return_sum_of_2_numbers(self):
        val = string_calculator('1,2')
        assert(val, 3)

    def should_return_sum_of_3_numbers(self):
        val = string_calculator('1,2,3')
        assert (val, 6)
