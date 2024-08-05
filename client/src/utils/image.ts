export function handleChamionImageName(_championName: string) {
  const championName = handleSwarm(_championName)
  switch (championName) {
    // ddragon use different case for only fiddlesticks, sadly
    case 'FiddleSticks':
      return 'Fiddlesticks'
    default:
      return championName
  }
}

function handleSwarm(championName: string): string {
  return championName.replace("Strawberry_", "")
}
