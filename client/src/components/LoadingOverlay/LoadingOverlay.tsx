import Spinner from './Spinner'

export default function LoadingOverlay() {
  return (
    <div className="fixed top-0 left-0 h-full w-full bg-white">
      <div className="w-full h-full flex justify-center items-center">
        <Spinner />
      </div>
    </div>
  )
}
