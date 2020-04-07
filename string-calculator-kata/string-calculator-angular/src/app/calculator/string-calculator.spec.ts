

describe('StringCalculator', () => {
  const stringCalculator = new StringCalculator();

  it('Should return 0 on blank string', () => {
    const val = stringCalculator.calculateValues('');
    expect(val).toEqual(0);
  });

  it('Should return correct sum with 2 numbers', () => {
    const val = stringCalculator.calculateValues('12,34');
    expect(val).toEqual(46);
  });

  it('Should return correct sum with 3 numbers', () => {
    const val = stringCalculator.calculateValues('1,2,3');
    expect(val).toEqual(6);
  });
});
