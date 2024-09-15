import { HOST, PORT, STATUS } from "./configs/constants.js";
import { aboutProject } from "./configs/data.js";
import { booksRoute } from "./routes/books.route.js";
import Hapi from "@hapi/hapi";
import dotenv from "dotenv";

dotenv.config();

async function main() {
  const server = Hapi.server({
    port: PORT,
    host: HOST,
  });

  server.route([
    {
      method: "GET",
      path: "/",
      handler: (_, h) => {
        return h
          .response({
            status: STATUS.success,
            message: "Sukses mendapatkan data About this project!",
            data: aboutProject,
          })
          .code(200);
      },
    },
    ...booksRoute,
  ]);

  await server.start();
  console.log(`Server sudah berjalan di port ${server.info.port}`);
}

main();
