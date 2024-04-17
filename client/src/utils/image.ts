export function handleChamionImageName(championName: string) {
  switch (championName) {
    // riot api and ddragon use different case for fiddlesticks, sadly
    case 'FiddleSticks':
      return 'Fiddlesticks'
    default:
      return championName
  }
}
