export interface IMatch {
  championName: string
  kills: number
  deaths: number
  assists: number
  win: boolean
  gameMode: string
  gameCreation: number
  participantList: Participant[]
}

export interface Participant {
  gameName: string
  tagLine: string
  championName: string
  kills: number
  deaths: number
  assists: number
  win: boolean
}
