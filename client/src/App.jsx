import './App.css'

const App = () => {
  return (
    <div class="flex flex-col items-center">
        <input type="file" accept="image/*" class="mb-10"></input>
        <button class="bg-blue-600 font-semibold text-white px-2 py-2 rounded">Upload gambar</button>
    </div>
  )
}

export default App
