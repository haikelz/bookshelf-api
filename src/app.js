import { HOST, PORT } from "./configs/constants.js";
import { booksRoute } from "./routes/books.route.js";
import Hapi from "@hapi/hapi";
import dotenv from "dotenv";

dotenv.config();

async function main() {
  const server = Hapi.server({
    port: PORT,
    host: HOST,
    routes: {
      cors: {
        origin: ["*"],
      },
    },
  });

  server.route(booksRoute);

  await server.start();
  console.log(`Server sudah berjalan di port ${server.info.port}`);
}

main();
