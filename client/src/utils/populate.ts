interface PopulateData {
  gameName: string
  tagLine: string
}

const POPULATE_DATA_STORAGE_KEY = 'populate'

// normally, I don't use try-catch inside a util function
// but I use it here because this one is an additional feature
// and I don't want to put several lines of code in main files

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

export function populateDate({
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
      // prefix #
      if (
        typeof data.tagLine === 'string' &&
        data.tagLine.length !== 0 &&
        !data.tagLine.includes('#')
      ) {
        setTagLine(`#${data.tagLine}`)
      }
    }
  } catch (err) {
    console.error(err)
  }
}
