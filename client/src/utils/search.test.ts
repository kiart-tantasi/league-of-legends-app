import { handleTagLine, validateSearchInputs, warnUser } from './search'

describe('handleTagLine', () => {
  test('Should remove hash symbol if has', () => {
    expect(handleTagLine({ tagLine: '#123' })).toBe('123')
    expect(handleTagLine({ tagLine: '123' })).toBe('123')
  })
})

describe('validateSearchInputs', () => {
  test('Should return true', () => {
    expect(
      validateSearchInputs({ gameName: 'ABC', tagLine: '#123' }),
    ).toBeTruthy()
  })
  test('Should return false', () => {
    expect(validateSearchInputs({ gameName: 'ABC', tagLine: '' })).toBeFalsy()
    expect(validateSearchInputs({ gameName: 'ABC', tagLine: null })).toBeFalsy()
  })
})

describe('warnUser', () => {
  const alertMock = jest.fn()
  window.alert = alertMock
  test('Should trigger alert', () => {
    warnUser('ERROR')
    expect(alertMock).toBeCalled()
  })
})
