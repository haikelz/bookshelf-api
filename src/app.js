import Hapi from "@hapi/hapi";
import dotenv from "dotenv";
import { HOST, PORT } from "./configs/constants.js";
import { booksRoute } from "./routes/books.route.js";

dotenv.config();

async function main() {
  const server = Hapi.server({
    port: PORT,
    host: HOST,
  });

  server.route(booksRoute);

  await server.start();
  console.log(`Server sudah berjalan di port ${server.info.port}`);
}

main();
