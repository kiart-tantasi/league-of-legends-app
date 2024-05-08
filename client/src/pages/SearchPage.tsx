import { FormEvent, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { handleTagLine, validateSearchInputs, warnUser } from '../utils/search'
import Layout from '../components/Layout'
import { searchPlaceholder } from '../configs/placeholder'
import { populateDate, savePopulateData } from '../utils/populate'

export default function SearchPage() {
  const [gameName, setGameName] = useState('')
  const [tagLine, setTagLine] = useState('')
  const navigate = useNavigate()

  const onSubmit = async (e: FormEvent) => {
    e.preventDefault()
    if (!validateSearchInputs({ gameName, tagLine })) {
      warnUser('กรอกข้อมูลไม่ครบ/ไม่ถูกต้อง')
      return
    }
    const handledTagLine = handleTagLine({ tagLine })
    savePopulateData({ gameName, tagLine: handledTagLine })
    navigate(`/match?gameName=${gameName}&tagLine=${handledTagLine}`)
  }

  useEffect(() => {
    populateDate({ setGameName, setTagLine })
  }, [])

  return (
    <Layout background="BLUE">
      <div className="flex flex-col h-full" data-testid="root-app">
        <form
          className="flex flex-col w-full items-center pt-[200px]"
          onSubmit={onSubmit}
        >
          <input
            value={gameName}
            onChange={(e) => setGameName(e.target.value)}
            type="text"
            name="gameName"
            placeholder={searchPlaceholder.gameName}
            className="w-[250px] mb-4 placeholder:text-[0.75rem]"
            autoFocus
          />
          <input
            value={tagLine}
            onChange={(e) => {
              setTagLine(e.target.value)
            }}
            type="text"
            name="tagLine"
            placeholder={searchPlaceholder.tagLine}
            className="w-[250px] mb-4 placeholder:text-[0.75rem]"
          />
          <button
            className="w-[250px] bg-white rounded p-1 text-[14px] font-bold
                      hover:bg-black hover:text-white"
            type="submit"
          >
            ดูประวัติการเล่น
          </button>
          <button
            onClick={() => {
              navigate('/match?gameName=เพชร&tagLine=ARAM')
            }}
            className="mt-[70px] w-[250px] bg-white rounded p-1 text-[10px] font-bold
                      hover:bg-black hover:text-white"
          >
            just want to try ?
          </button>
        </form>
        <div className="text-center mt-[200px]">
          โดย&nbsp;
          <a
            href="https://www.petchblog.net"
            className="font-bold"
            target="_blank"
            rel="noopener noreferrer"
          >
            petchblog.net
          </a>
        </div>
      </div>
    </Layout>
  )
}
