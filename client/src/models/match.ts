export interface IMatch {
  championName: string
  kills: number
  deaths: number
  assists: number
  win: boolean
  gameMode: string
  gameCreation: number
  participants: Participant[]
}

export interface Participant {
  gameName: string
  championName: string
  kills: number
  deaths: number
  assists: number
  win: boolean
  puuid: string
}
