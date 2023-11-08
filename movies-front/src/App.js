import Container from "./components/Container"
import './App.css';
import { Tabs } from 'flowbite-react';
import { TbMovie, TbChairDirector } from 'react-icons/tb'
import ListMovies from './components/ListMovies'
import ListDirectors from "./components/ListDirectors";

function App() {
  return (
    <Container title="Movies App" >
      <Tabs.Group
        aria-label="Tabs with underline"
        style="underline"
      >
        <Tabs.Item
          active
          icon={TbMovie}
          title="Movies"
        >
          <ListMovies />
        </Tabs.Item>
        <Tabs.Item
          icon={TbChairDirector}
          title="Directors"
        >
          <ListDirectors />
        </Tabs.Item>

      </Tabs.Group>
    </Container >
  );
}

export default App;
