//The search widget variables
const form = document.querySelector(".form");
const search = document.querySelector(".search");
const submit = document.querySelector(".submit");

//variables to display movie
const content = document.getElementById("movies");

//API urls from the movie db
const API_URL =
  "https://api.themoviedb.org/3/discover/movie?sort_by=popularity.desc&api_key=95c4f87fa75a8ad6973f481f2352760a&page=1";

const IMG_PATH = "https://image.tmdb.org/t/p/w1280/";

const GENRE_PATH =
  "https://api.themoviedb.org/3/genre/movie/list?api_key=95c4f87fa75a8ad6973f481f2352760a&language=en-US";

const SEARCH_API =
  "https://api.themoviedb.org/3/search/movie?api_key=95c4f87fa75a8ad6973f481f2352760a&query='";


//Getting initial movies
getMovies(API_URL, GENRE_PATH);

async function getMovies(url, gen_url) {
  //for movie data
  const res = await fetch(url);
  const data = await res.json();

  //for genre data
  const gen = await fetch(gen_url);
  const genre = await gen.json();

  results = data.results;
  genres = genre.genres;

  const new_movie = results.map(({ genre_ids, ...rest }) => ({
    ...rest,
    genre_ids: genre_ids.map(
      (id) => genres.find((genre) => genre.id === id).name
    ),
  }));

  showMovies(new_movie);
}

function showMovies(movies) {
  content.innerHTML = "";

  movies.forEach((movie) => {
    const { title, poster_path, vote_average, genre_ids, original_language } = movie;

    const the_movie = document.createElement("div");
    the_movie.classList.add("movie-info");

    the_movie.innerHTML = `
        <div class="poster">
            <h2>${title}</h2>
          
            <img src="${IMG_PATH + poster_path}" alt="poster">
            <h4>${genre_ids.join(", ")}</h4>
        </div>
        <div class="rating ${getClassByRate(vote_average)}">
            Rating:
            <span>${vote_average}</span>
        </div>
        <div class="trailer">
        P
        </div>
        <div class = "language" >
            <h3>${original_language}</h3>
        </div>
        `;
    content.appendChild(the_movie);
  });
}

function getClassByRate(vote) {
  if (vote >= 8) {
    return "green";
  } else if (vote >= 5) {
    return "orange";
  } else {
    return "red";
  }
}

form.addEventListener("submit", (e) => {
  e.preventDefault();

  const searchTerm = search.value;

  if (searchTerm && searchTerm !== "") {
    getMovies(SEARCH_API + searchTerm);

    search.value = "";
  } else {
    window.location.reload();
  }
});
