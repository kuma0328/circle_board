import { useState } from "react"
import { CircleGenreInput } from "../components/circle/CircleGenreInput"
import { CircleIconInput } from "../components/circle/CircleIconInput"
import { CircleInstagramInput } from "../components/circle/CircleInstagramInput"
import { CircleLocationInput } from "../components/circle/CircleLocationInput"
import { CircleMemberInput } from "../components/circle/CircleMemberInput"
import { CircleNameInput } from "../components/circle/CircleNameInput"
import { CircleTwitterInput } from "../components/circle/CircleTwitterInput"

export const CircleForm = () => {
  const [circleName, setCircleName] = useState("")
  const [circleIcon, setCircleIcon] = useState(null)
  const [circleGenre, setCircleGenre] = useState("")
  const [circleLocation, setCircleLocation] = useState("")
  const [circleMember, setCircleMember] = useState(0)
  const [circleTwitter, setCircleTwitter] = useState("")
  const [circleInstagram, setCircleInstagram] = useState("")

  const handleSubmit = (e) => {
    e.preventDefault()
    const formData = {
      circleName,
      circleIcon,
      circleGenre,
      circleLocation,
      circleMember,
      circleTwitter,
      circleInstagram,
    }
    console.log(formData)
    fetch("localhost:8081/circle/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      })
      .then((res) => res.json())
      .then((data) => console.log(data))
  }

  return (
    <div className="flex flex-col text-xl">
      <form onSubmit={handleSubmit}>
        <CircleNameInput circleName={circleName} setCircleName={setCircleName}/>
        <CircleIconInput circleIcon={circleIcon} setCircleIcon={setCircleIcon}/>
        <CircleGenreInput circleGenre={circleGenre} setCircleGenre={setCircleGenre}/>
        <CircleLocationInput circleLocation={circleLocation} setCircleLocation={setCircleLocation}/>
        <CircleMemberInput circleMember={circleMember} setCircleMember={setCircleMember}/>
        <CircleTwitterInput circleTwitter={circleTwitter} setCircleTwitter={setCircleTwitter}/>
        <CircleInstagramInput circleInstagram={circleInstagram} setCircleInstagram={setCircleInstagram} />
        <button
          type="submit"
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        >
          送信
        </button>
      </form>
    </div>
  )
}
