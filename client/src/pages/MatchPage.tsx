import { FormEvent, useContext, useEffect, useState } from 'react'
import MatchContext from '../contexts/MatchContext'
import { Link, useSearchParams } from 'react-router-dom'
import {
  handleTagLine,
  validateSearchInputs,
  warnUser,
} from './../utils/search'
import Layout from '../components/Layout'
import { IMatch, Participant } from '../models/match'
import getMatchDetailList from '../api/getMatchDetailList'
import { Size } from '../constants/common'
import { searchPlaceholder } from '../configs/placeholder'
import { handleChamionImageName as handleImageName } from '../utils/image'
import LoadingOverlay from '../components/LoadingOverlay/LoadingOverlay'
import { ddragonConfig } from '../configs/ddragon'

const PRIORITIZED_CARD_AMOUNT = 4

export default function MatchPage() {
  const [searchParams, setSearchParams] = useSearchParams()
  const paramGameName = searchParams.get('gameName')
  const paramTagLine = searchParams.get('tagLine')
  const [gameName, setGameName] = useState('')
  const [tagLine, setTagLine] = useState('')
  const { matches, setMatches } = useContext(MatchContext)
  const [isLoading, setIsLoading] = useState(true)
  const [imageLoadedCounter, setImageLoadedCounter] = useState(0)
  // isLoadingV2 considers loading as done when 4 champion-images are also loaded
  const isLoadingV2 =
    isLoading ||
    imageLoadedCounter < Math.min(PRIORITIZED_CARD_AMOUNT, matches.length)

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
        setImageLoadedCounter(0)
        const { matchDetailList, status } = await getMatchDetailList({
          paramGameName: paramGameName || '',
          paramTagLine: paramTagLine || '',
        })
        if (status !== 200) {
          warnUser('ไม่พบผู้เล่น/เซิร์ฟเวอร์กำลังทำงานหนัก กรุณาลองใหม่')
        }
        setMatches(matchDetailList)
      } catch (e) {
        setMatches([])
        console.error(e)
      } finally {
        setIsLoading(false)
      }
    })()
  }, [paramGameName, paramTagLine])

  const onSubmit = (e: FormEvent) => {
    e.preventDefault()
    if ([gameName, tagLine].some((e) => !e)) {
      warnUser('กรอกข้อมูลไม่ครบ/ไม่ถูกต้อง')
      return
    }
    // reload the page if gameName and tagLine are the same as current ones
    if (paramGameName === gameName && paramTagLine === tagLine) {
      window.location.reload()
      return
    }
    setSearchParams({
      gameName,
      tagLine: handleTagLine(tagLine),
    })
  }

  return (
    <>
      {isLoadingV2 && <LoadingOverlay />}
      <Layout>
        <div className="flex flex-col justify-center pt-2 w-full max-w-[600px]">
          <div className="flex justify-between mb-4">
            <Link className="p-2 w-fit h-fit mb-4 border" type="button" to="/">
              หน้าแรก
            </Link>
            <form className="flex flex-col w-[150px]" onSubmit={onSubmit}>
              <input
                value={gameName}
                onChange={(e) => setGameName(e.target.value)}
                placeholder={searchPlaceholder.gameName}
                type="text"
                className="text-right mb-2"
              />
              <input
                value={tagLine}
                onChange={(e) => setTagLine(e.target.value)}
                placeholder={searchPlaceholder.tagLine}
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
          {matches.map((match, index) => {
            const shouldLazy = index > PRIORITIZED_CARD_AMOUNT - 1
            const countImageLoaded = () => {
              setImageLoadedCounter((prev) => prev + 1)
            }
            return (
              <MatchCard
                match={match}
                key={`match-detail-${index}`}
                shouldLazy={shouldLazy}
                onImageLoaded={shouldLazy ? null : countImageLoaded}
              />
            )
          })}
        </div>
      </Layout>
    </>
  )
}

function MatchCard({
  match,
  shouldLazy,
  onImageLoaded,
}: {
  match: IMatch
  shouldLazy: boolean
  onImageLoaded: (() => void) | null
}) {
  const backgroundColor = match.win ? 'bg-blue-100' : 'bg-red-100'
  const [isOpen, setIsOpen] = useState(false)
  return (
    <>
      <div
        className={`p-2 mb-1 ${backgroundColor} ${
          isOpen ? 'border-b border-b-gray-300' : ''
        }`}
        onClick={() => setIsOpen((prev) => !prev)}
      >
        <div className="flex flex-row justify-between">
          <div className="flex">
            <ChampionImage
              championName={match.championName}
              size={Size.BIG}
              shouldLazy={shouldLazy}
              onImageLoaded={onImageLoaded}
            />
            <div className="flex ml-1">
              {match.itemIds.map((itemId, index) => (
                <ItemImage
                  itemId={itemId}
                  size={Size.BIG}
                  shouldLazy={shouldLazy}
                  key={`match-card-item-id-${itemId}-${index}`}
                />
              ))}
            </div>
          </div>
          <p>{`${match.kills}/${match.deaths}/${match.assists}`}</p>
        </div>
        <p className="text-[0.7rem] mt-6">{match.gameMode}</p>
        <div className="flex justify-between">
          <p className="text-[0.55rem]">
            {new Date(match.gameCreation).toLocaleDateString('pt-PT')}
          </p>
          <button className="font-bold text-[0.6rem]">
            {isOpen ? 'ย่อ' : 'ขยาย'}
          </button>
        </div>
      </div>
      {isOpen && (
        <div className="pr-4 mb-4">
          {match.participantList.map((parti, index) => (
            <ParticipantCard parti={parti} key={`participant-${index}`} />
          ))}
        </div>
      )}
    </>
  )
}

function ParticipantCard({ parti }: { parti: Participant }) {
  return (
    <a
      className={`flex justify-between px-2 py-1 border-b ${
        parti.win ? 'bg-blue-100' : 'bg-red-100'
      }`}
      href={`/match?gameName=${parti.gameName}&tagLine=${parti.tagLine}`}
      target="_blank"
      rel="noopener noreferrer"
    >
      <div>
        <div>
          <div className="flex">
            <ChampionImage
              championName={parti.championName}
              size={Size.SMALL}
              shouldLazy={false}
            />
            <div className="flex ml-1">
              {parti.itemIds.map((itemId, index) => (
                <ItemImage
                  itemId={itemId}
                  size={Size.SMALL}
                  shouldLazy={false}
                  key={`participant-card-${parti.gameName}-item-id-${itemId}-${index}`}
                />
              ))}
            </div>
          </div>
          <p className="text-[0.7rem] mt-1">{parti.gameName}</p>
        </div>
      </div>
      <div>
        <p>{`${parti.kills}/${[parti.deaths]}/${parti.assists}`}</p>
      </div>
    </a>
  )
}

function ChampionImage({
  championName,
  size,
  shouldLazy,
  onImageLoaded,
}: {
  championName: string
  size: Size
  shouldLazy: boolean
  onImageLoaded?: (() => void) | null
}) {
  const [isError, setIsError] = useState(false)
  const widthHeightClass =
    size === Size.BIG ? `w-12 h-12 md:w-12 md:h-12` : 'w-10 h-10 md:w-9 md:h-9'
  if (isError) {
    return (
      <div data-champion-name={championName} className={widthHeightClass} />
    )
  }
  return (
    <img
      className={widthHeightClass}
      src={`https://ddragon.leagueoflegends.com/cdn/${
        ddragonConfig.version
      }/img/champion/${handleImageName(championName)}.png`}
      alt={`${championName}`}
      onError={() => {
        setIsError(true)
        !shouldLazy && onImageLoaded && onImageLoaded()
      }}
      onLoad={() => {
        !shouldLazy && onImageLoaded && onImageLoaded()
      }}
      loading={shouldLazy ? 'lazy' : 'eager'}
    />
  )
}

function ItemImage({
  itemId,
  size,
  shouldLazy,
}: {
  itemId: number
  size: Size
  shouldLazy: boolean
}) {
  const [isError, setIsError] = useState(false)
  const widthHeightClass =
    size === Size.BIG ? 'w-9 h-9 md:w-10 md:h-10' : 'w-8 h-8 md:w-8 md:h-8'
  if (isError) {
    return <div data-item-id={itemId} className={widthHeightClass} />
  }
  return (
    <img
      className={widthHeightClass}
      src={`https://ddragon.leagueoflegends.com/cdn/${ddragonConfig.version}/img/item/${itemId}.png`}
      alt={`league of legends item id ${itemId}`}
      onError={() => setIsError(true)}
      loading={shouldLazy ? 'lazy' : 'eager'}
    />
  )
}
