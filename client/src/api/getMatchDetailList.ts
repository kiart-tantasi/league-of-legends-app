import matchDetailListMock from '../mocks/matchDetailsList'
import { handleTagLine } from '../utils/search'

export default async function getMatchDetailList({
  paramGameName,
  paramTagLine,
}: {
  paramGameName: string
  paramTagLine: string
}) {
  // ========= DEV ========= //
  const REACT_APP_IS_MOCK = false // in case you want to mock results so you don't have to connect to backend api
  if (REACT_APP_IS_MOCK) {
    return {
      status: 200,
      matchDetailList: matchDetailListMock,
    }
  }
  // ========= DEV ========= //

  const response = await fetch(
    `/api/v1/matches?gameName=${paramGameName}&tagLine=${handleTagLine(
      paramTagLine,
    )}`,
  )
  let matchDetailList = []
  if (response.status === 200) {
    const json = await response.json()
    matchDetailList = json.matchDetailList
  }
  return {
    status: response.status,
    matchDetailList,
  }
}
