import { use } from "react"
import { useState } from "react"

const App = () => {
  const [image, setImage] = useState()
  const [btnText, setBtnText] = useState("Convert image")
  const [selectedFile, setSelectedFile] = useState(null)
  const [isNotReadyToConvert, setIsNotReadyToConvert] = useState(true)
  const [isConverted, setIsConverted] = useState()

  if (isNotReadyToConvert == false) {
    document.getElementById("action-btn").classList.remove("cursor-not-allowed")
    document.getElementById("action-btn").classList.remove("bg-slate-400")
    document.getElementById("action-btn").classList.add("bg-blue-400")
  }

  const handleChange = (e) => {
    const file = e.target.files[0]

    setImage(URL.createObjectURL(file))
    setSelectedFile(file)

    setIsNotReadyToConvert(false)
    setBtnText("Convert image")
    setIsConverted(false)
  }

  const handleClick = () => {
    if (!isConverted) {
      convertImage()
    } else {
      downloadImage()
    }
  }

  const downloadImage = () => {
    const link = document.createElement("a")
    link.href = image
    link.download = "result.png"

    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  }

  const convertImage = async() => {
    const formData = new FormData()
    formData.append("file", selectedFile)
    
    try {
      const response = await fetch("http://localhost:8080/upload", {
        method: "POST",
        body: formData,
      })

      if (!response.ok) throw new Error("Gagal memproses gambar")

      const blob = await response.blob()
      const imgUrl = URL.createObjectURL(blob)
  
      setImage(imgUrl)
      setBtnText("Download image")
      setIsConverted(true)

    } catch(error) {
        console.error(error)
    }

  }

  return (
    <>
      <div class="flex items-center justify-center w-95 h-[45rem] px-6 pt-4 pb-4">
        {isNotReadyToConvert && (
          <label id="dropzone" for="dropzone-file" class="flex flex-col items-center justify-center w-full h-full border-2 border-gray-300 border-dashed rounded-lg cursor-pointer bg-gray-50 dark:hover:bg-gray-800 dark:bg-gray-700 hover:bg-gray-100 dark:border-gray-600 dark:hover:border-gray-500 dark:hover:bg-gray-600">
            <div class="flex flex-col items-center justify-center pt-5 pb-6">
                <svg class="w-8 h-8 mb-4 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 16">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 13h3a3 3 0 0 0 0-6h-.025A5.56 5.56 0 0 0 16 6.5 5.5 5.5 0 0 0 5.207 5.021C5.137 5.017 5.071 5 5 5a4 4 0 0 0 0 8h2.167M10 15V6m0 0L8 8m2-2 2 2"/>
                </svg>
                <p class="mb-2 text-md text-gray-500 dark:text-gray-400"><span class="font-semibold">Click to upload</span> or drag and drop</p>
                <p id="file-info" class="text-sm text-gray-500 dark:text-gray-400">PNG FILE</p>
            </div>
          </label>
        )}

        {!isNotReadyToConvert && (
          <label id="preview-of-input" for="dropzone-file" class="flex flex-col items-center justify-center w-full h-full border-2 border-gray-300 border-dashed rounded-lg cursor-pointer bg-gray-50 dark:hover:bg-gray-800 dark:bg-gray-700 hover:bg-gray-100 dark:border-gray-600 dark:hover:border-gray-500 dark:hover:bg-gray-600">
            <img id="preview" src={image} class="h-full"/>
          </label>
        )}

        <input id="dropzone-file" type="file" accept="image/png" class="hidden" onChange={handleChange} />

      </div> 
      
      <div class="flex flex-col items-center">
          <button id="action-btn" class="bg-slate-400 font-semibold text-white px-2 py-2 rounded cursor-not-allowed" onClick={handleClick} disabled={isNotReadyToConvert}>{btnText}</button>
      </div>
    </>
  )
}

export default App
