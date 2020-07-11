package main

const publicBegining = `
<!doctype html>
<html lang="en">
	<head>
		<meta name="description" content="Strona osobista Roberta Bendun, zawierająca rozmaite narzędzia webowe i desktopowe, ciekawe informacje i wartościowe wpisy." />
		<meta name="color-scheme" content="dark light">
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<link rel="stylesheet" href="https://bendun.cc/common.css">
		<link rel="stylesheet" href="https://bendun.cc/common-mobile.css">
		<script type="module" src="https://bendun.cc/dark-mode.js"></script>
		<title>Public short links</title>
		<style>
		table, tr {
			width: 100%;
			text-align: center;
		}
		</style>
	</head>
<body>
<header>
    <h1>Robert Bendun</h1>
    <nav>
      <ul>
        <li><a href="https://bendun.cc/">Główna</a></li>
        <li><a target="_blank" href="https://github.com/RobertBendun">Github</a></li>
        <li><a href="https://bendun.cc/wiki">Wiki</a></li>
      </ul>
    </nav>
    <button id="change-mode">🌃</button>
  </header>
  <main>
		<h2>List of all public links</h1>
		<table>
			<tr>
				<th>Path</th>
				<th>Target URL</th>
			</tr>
`

const publicEnding = `
	</table>
	</main>
</body>
</html>
`
