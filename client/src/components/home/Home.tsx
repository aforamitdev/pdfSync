import { accounts, mails } from "./data"
import { Mail } from "./components/mail"


export default function Home() {

  return (
    <>
  
      <div className="hidden flex-col md:flex">
        <Mail
          accounts={accounts}
          mails={mails}

          navCollapsedSize={4}
        />
      </div>
    </>
  )
}
