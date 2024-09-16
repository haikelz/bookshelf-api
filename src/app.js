import { HOST, PORT } from "./configs/constants.js";
import { AboutProject, Status } from "./entities/books.entity.js";
import { booksRoute } from "./routes/books.route.js";
import Hapi from "@hapi/hapi";
import dotenv from "dotenv";

dotenv.config();

async function main() {
  const server = Hapi.server({
    port: PORT,
    host: HOST,
  });

  const status = new Status();

  const aboutProject = new AboutProject();

  server.route([
    {
      method: "GET",
      path: "/",
      handler: (_, h) => {
        return h
          .response({
            status: status.success,
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
