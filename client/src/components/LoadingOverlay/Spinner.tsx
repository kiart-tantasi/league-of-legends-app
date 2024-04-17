import styles from './Spinner.module.css'

export default function Spinner() {
  return (
    <div
      className={` ${styles.spinnerAnimation}
                    border-4 border-blue-500 border-b-transparent
                    w-10 h-10 rounded-full`}
    />
  )
}
