import { FormEvent, useContext, useEffect, useState } from 'react'
import MatchContext from '../contexts/MatchContext'
import { Link, useSearchParams } from 'react-router-dom'
import { validateSearchInputs, warnUser } from './../utils/search'
import Layout from '../components/Layout'
import { IMatch, Participant } from '../models/match'
import getMatchDetailList from '../api/getMatchDetailList'

export default function MatchPage() {
  const [searchParams, setSearchParams] = useSearchParams()
  const paramGameName = searchParams.get('gameName')
  const paramTagLine = searchParams.get('tagLine')
  const [gameName, setGameName] = useState('')
  const [tagLine, setTagLine] = useState('')
  const { matches, setMatches } = useContext(MatchContext)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    if (
      !validateSearchInputs({ gameName: paramGameName, tagLine: paramTagLine })
    ) {
      warnUser('ไม่พบชื่อ/tagline')
      setIsLoading(false)
      return
    }

    ;(async () => {
      try {
        setIsLoading(true)
        const { matchDetailList, status } = await getMatchDetailList({
          paramGameName: paramGameName || '',
          paramTagLine: paramTagLine || '',
        })
        if (status !== 200) {
          warnUser('ไม่พบ กรุณาลองใหม่')
        }
        setMatches(matchDetailList)
      } catch (e) {
        console.error(e)
      } finally {
        setIsLoading(false)
      }
    })()
  }, [paramGameName, paramTagLine, setMatches])

  const onSubmit = (e: FormEvent) => {
    e.preventDefault()
    if (
      [gameName, tagLine].some((e) => typeof e !== 'string' || e.length === 0)
    ) {
      warnUser('กรอกข้อมูลไม่ครบ/ไม่ถูกต้อง')
      return
    }
    setSearchParams({
      gameName,
      tagLine,
    })
  }

  if (isLoading) {
    return (
      <Layout>
        <div className="text-center pt-[200px]">กำลังโหลด...</div>
      </Layout>
    )
  }
  return (
    <Layout>
      <div className="flex flex-col justify-center pt-2 w-full max-w-[600px]">
        <div className="flex justify-between mb-4">
          <Link className="p-2 w-fit h-fit mb-4 border" type="button" to="/">
            กลับ
          </Link>
          <form className="flex flex-col w-[150px]" onSubmit={onSubmit}>
            <input
              value={gameName}
              onChange={(e) => setGameName(e.target.value)}
              placeholder="ชื่อในเกม"
              type="text"
              className="text-right mb-2"
            />
            <input
              value={tagLine}
              onChange={(e) => setTagLine(e.target.value)}
              placeholder="#1234"
              type="text"
              className="text-right mb-2"
            />
            <button type="submit" className="border">
              ดู match history
            </button>
          </form>
        </div>
        <div className="p-2 bg-gray-200">
          {paramGameName} #{paramTagLine}
        </div>
        {matches.map((match, index) => (
          <MatchCard match={match} key={`match-detail-${index}`} />
        ))}
      </div>
    </Layout>
  )
}

function MatchCard({ match }: { match: IMatch }) {
  const backgroundColor = match.win ? 'bg-blue-100' : 'bg-red-100'
  const [isOpen, setIsOpen] = useState(false)
  return (
    <div className="mb-1">
      <div
        className={` p-2 ${backgroundColor}`}
        onClick={() => setIsOpen((prev) => !prev)}
      >
        <div className="flex flex-row justify-between">
          <p>{match.championName}</p>
          <p>{`${match.kills}/${match.deaths}/${match.assists}`}</p>
        </div>
        <p className="text-[0.7rem]">{match.gameMode}</p>
        <div className="flex justify-between">
          <p className="text-[0.7rem]">
            {new Date(match.gameCreation).toLocaleDateString('pt-PT')}
          </p>
          <button className="font-bold text-[0.75rem]">
            {isOpen ? 'ปิด' : 'ดูข้อมูล'}
          </button>
        </div>
      </div>
      {isOpen && (
        <div>
          {match.participantList.map((parti, index) => (
            <ParticipantCard parti={parti} key={`participant-${index}`} />
          ))}
        </div>
      )}
    </div>
  )
}

function ParticipantCard({ parti }: { parti: Participant }) {
  return (
    <a
      className="flex justify-between p-2 bg-gray-100 border-b"
      href={`/match?gameName=${parti.gameName}&tagLine=${parti.tagLine}`}
      target="_blank"
      rel="noopener noreferrer"
    >
      <div>
        <p>{parti.gameName}</p>
        <p className="text-[0.75rem]">{parti.championName}</p>
      </div>
      <div>
        <p>{`${parti.kills}/${[parti.deaths]}/${parti.assists}`}</p>
      </div>
    </a>
  )
}
