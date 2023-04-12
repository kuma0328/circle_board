import {useNavigate} from 'react-router-dom'

export const FormButton = () => {
  const navigate = useNavigate()

  return (
    <>
      <button onClick={() => {navigate("/form")}}>サークルを登録する</button>
    </>
  )
}

