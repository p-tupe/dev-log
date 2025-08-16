/***
 * DESCRIPTION:
 *   A small scraper to fetch chapters of a story from RoyalRoad.com.
 *   Give it the url of a chapter and it keeps pulls all chapters until the end!
 *
 * USAGE: deno run scrape.ts <royalroad.com/fiction/x/y/chapter/x/y [text | html (default)]
 *
 * OUTPUT: story.[txt | html]
 */

import cheerio from "https://dev.jspm.io/cheerio";

const FILE_FORMAT = Deno.args[1] === "text" ? "text" : "html";
const FILE_PATH = `./story.${FILE_FORMAT === "text" ? "txt" : "html"}`;

async function* fetchChapter() {
  let nextChapterLink = Deno.args[0];

  while (nextChapterLink) {
    console.log("Fetching", nextChapterLink);

    const res = await fetch(nextChapterLink);
    const html = await res.text();
    const $ = cheerio.load(html);

    const chapterContent = $("div.chapter-content");
    yield chapterContent[FILE_FORMAT]();

    nextChapterLink = $('a.btn.btn-primary:contains("Next")')?.attr("href");
    if (nextChapterLink)
      nextChapterLink = "https://www.royalroad.com" + nextChapterLink;
  }
}

async function main() {
  const outputStream = await (
    await Deno.open(FILE_PATH, {
      create: true,
      append: true,
    })
  ).writable.getWriter();

  const encoder = new TextEncoder();

  for await (const content of fetchChapter()) {
    const encodedContent = encoder.encode("\n\n---\n\n" + content);
    outputStream.write(encodedContent);
  }

  await outputStream.close();

  console.log("All Done!");
}

main();
