import { FormEvent, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { validateSearchInputs, warnUser } from '../utils/search'
import Layout from '../components/Layout'

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
    navigate(`/match?gameName=${gameName}&tagLine=${tagLine.replace('#', '')}`)
  }

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
            placeholder="ชื่อในเกม"
            className="w-[250px] mb-4 placeholder:text-[0.75rem]"
          />
          <input
            value={tagLine}
            onChange={(e) => {
              setTagLine(e.target.value)
            }}
            type="text"
            name="tagLine"
            placeholder="#1234"
            className="w-[250px] mb-4 placeholder:text-[0.75rem]"
          />
          <button className="bg-white rounded p-1 text-[14px]" type="submit">
            ดู match history
          </button>
        </form>
        <div className="text-center mt-[200px]">
          by&nbsp;
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
