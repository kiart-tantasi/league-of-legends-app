// normally, I don't use try-catch inside a util function
// but I use it here because this one is an additional feature
// and I don't want to put several lines of code in main files

interface PopulateData {
  gameName: string
  tagLine: string
}

const POPULATE_DATA_STORAGE_KEY = 'populate'

export function savePopulateData({
  gameName,
  tagLine,
}: {
  gameName: string
  tagLine: string
}): void {
  try {
    const data: PopulateData = {
      gameName,
      tagLine,
    }
    localStorage.setItem(POPULATE_DATA_STORAGE_KEY, JSON.stringify(data))
  } catch (err) {
    console.error(err)
  }
}

export function populateData({
  setGameName,
  setTagLine,
}: {
  setGameName: React.Dispatch<React.SetStateAction<string>>
  setTagLine: React.Dispatch<React.SetStateAction<string>>
}) {
  try {
    const str = localStorage.getItem(POPULATE_DATA_STORAGE_KEY)
    const data: PopulateData = JSON.parse(str || '')
    if (!!data) {
      setGameName(data.gameName || '')
      setTagLine(withHashtag(data.tagLine))
    }
  } catch (err) {
    console.error(err)
  }
}

// prefix hashtag if not already included
function withHashtag(tagLine: string | undefined): string {
  if (typeof tagLine === 'string' && tagLine.length !== 0) {
    return `${tagLine.includes('#') ? '' : '#'}${tagLine}`
  }
  return ''
}
