const form = document.querySelector("form");
const textarea = document.querySelector("textarea");
const btnScript1 = document.querySelector(".btn-script1");
const btnScript2 = document.querySelector(".btn-script2");

const sendRequest = async (url) => {
  try {
    const response = await fetch(url);

    if (response.ok) return response.text();

    throw new Error(`Request failed with status ${response.status}`);
  } catch (error) {
    console.error(error);
  }
};

form.addEventListener("submit", (event) => {
  event.preventDefault();
  const params = textarea.value.trim().split("\n");
  const queryParams = params
    .map((param) => encodeURIComponent(param))
    .join(",");
  const url = `http://127.0.0.1:17000/?cmd=${queryParams}`;

  sendRequest(url);
});

btnScript1.addEventListener("click", () => {
  const url =
    "http://127.0.0.1:17000/?cmd=green,bgrect%200.25%200.25%200.75%200.75,update";

  sendRequest(url);
});

btnScript2.addEventListener("click", () => {
  const urlToDraw =
    "http://127.0.0.1:17000/?cmd=white,figure%200.05%200.05,update";
  const urlToMove = "http://127.0.0.1:17000/?cmd=move%200.05%200.05,update";

  sendRequest(urlToDraw);

  for (let i = 0; i < 21; i++) {
    setTimeout(() => {
      sendRequest(urlToMove);
    }, (i + 1) * 1000);
  }
});