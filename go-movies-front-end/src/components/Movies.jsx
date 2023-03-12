import React, { useEffect, useState } from 'react'
import { Link } from 'react-router-dom';

const Movies = () => {
const [movies, setMovies] = useState([]);

useEffect(()=>{
  // let moviesList = [
  //     {
  //         id:1,
  //         title:"Highlander",
  //         release_date:'1986-03-07',
  //         runtime:116,
  //         mpaa_rating:"R",
  //         description:"Some long description"
  //     },
  //     {
  //         id:2,
  //         title:"Raiders of the Lost Ark",
  //         release_date:'1981-06-12',
  //         runtime:115,
  //         mpaa_rating:"PG-13",
  //         description:"Some long description"
  //     }
  // ];
  //Get the movies list from the back end

  async function getData() {
    const headers = new Headers();
    headers.append("Content-Type", "application/json");

    const requestOptions = {
      method: "GET",
      headers: headers,
    };

    try {
      var data = await fetch(`http://localhost:8080/movies`, requestOptions);
      var response = await data.json();
      setMovies(response);
    } catch (err) {
      console.log(err);
    }
  }

  getData();
},[])

  return (
    <div>
      <h2>Movies</h2>
      <hr />
      <table className="table table-striped table-hover">
        <thead>
          <tr>
            <th>Movie</th>
            <th>Release Date</th>
            <th>Rating</th>
          </tr>
        </thead>
        <tbody>
          {movies.map((m) => (
            <tr key={m.id}>
              <td>
                <Link to={`/movies/${m.id}`} >
                    {m.title}
                </Link>
              </td>
              <td>{m.release_date}</td>
              <td>{m.mpaa_rating}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default Movies