-- users table
DROP TABLE IF EXISTS users;

DROP TABLE IF EXISTS products;

DROP TABLE IF EXISTS orders;

CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  name text not null,
  username text not null,
  email text not null,
  password text not null,
  privilege TEXT CHECK(privilege IN ('user', 'visor', 'admin'))
);

CREATE TABLE products (
  id INTEGER PRIMARY KEY,
  name text not null,
  categories text DEFAULT "no category" NOT NULL,
  price decimal(5.2) not null,
  discount decimal(2.2) DEFAULT 0.0,
  description text,
  images text
);

INSERT INTO
  users (name, username, email, password, privilege)
VALUES
  (
    'test user',
    'testUser',
    'test_user@domain.com',
    'user@123',
    'user'
  ),
  (
    'test visor',
    'testVisor',
    'test_visor@domain.com',
    'visor@123',
    'visor'
  ),
  (
    'test admin',
    'testAdmin',
    'test_admin@domain.com',
    'admin@123',
    'admin'
  );

insert into
  products (
    name,
    categories,
    price,
    discount,
    description,
    images
  )
VALUES
  (
    "consequatur",
    "cat1",
    409.95,
    25,
    "Ullam illo non quisquam dicta tempora ad possimus repellendus.Eum et aperiam fugiat consequatur a ex.Placeat omnis voluptatem sit est.Corporis delectus vero doloremque.,",
    "http://placeimg.com/640/480/cats,http://placeimg.com/640/480/food"
  ),
  (
    "officiis",
    "cat2",
    514.21,
    12.5,
    "Aspernatur qui quia temporibus et facilis.Ut modi et eos.Error commodi laboriosam et consectetur et dolorem odio.Ea voluptates quia blanditiis dolores ut.,",
    "http://placeimg.com/640/480/cats,http://placeimg.com/640/480/food"
  ),
  (
    "odio",
    "cat3",
    983.60,
    40,
    "Dicta praesentium optio assumenda labore sed.Hic sequi eveniet soluta vitae.Recusandae dignissimos ab accusantium non distinctio nulla.Impedit modi ut et nemo.,",
    "http://placeimg.com/640/480/cats,http://placeimg.com/640/480/food"
  ),
  (
    "dolorum",
    "cat2",
    652.52,
    25,
    "Facilis optio neque minima nam ipsam.Atque a voluptate.Placeat enim molestiae earum.Repellat ut dicta assumenda commodi nesciunt suscipit sunt exercitationem.,",
    "http://placeimg.com/640/480/cats,http://placeimg.com/640/480/food"
  ),
  (
    "dolorum",
    "cat3",
    526.18,
    80,
    "Sed qui et.Aut et hic molestias culpa.Quis voluptates perspiciatis non incidunt ut dolorem deleniti maxime veniam.,",
    "http://placeimg.com/640/480/cats,http://placeimg.com/640/480/food"
  ),
  (
    "placeat",
    "cat1",
    719.25,
    25,
    "Voluptas expedita quia fuga sed nihil animi commodi facilis itaque.Sed aut eos sint voluptas.Dolor et consequatur.,",
    "http://placeimg.com/640/480/cats,http://placeimg.com/640/480/food"
  ),
  (
    "culpa",
    "cat1",
    775.63,
    10,
    "Alias eligendi est eum maxime.Nulla voluptas pariatur non.Voluptatem ipsam dolor.Totam velit vitae.Facere rerum et aut sapiente quaerat qui qui corporis.,",
    "http://placeimg.com/640/480/cats,http://placeimg.com/640/480/food"
  ),
  (
    "repellendus",
    "cat3",
    676.01,
    25,
    "Quibusdam porro corrupti laudantium dolores reprehenderit ducimus rem sed.Quia aliquid fugit.Aut non debitis ad vitae id voluptatem.Sapiente voluptatem cum.In sed voluptatem beatae.Qui repellendus provident eum.,",
    "http://placeimg.com/640/480/cats,http://placeimg.com/640/480/food"
  );