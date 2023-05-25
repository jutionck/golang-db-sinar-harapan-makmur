--
-- PostgreSQL database dump
--

-- Dumped from database version 14.6 (Homebrew)
-- Dumped by pg_dump version 14.6 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

ALTER TABLE IF EXISTS ONLY public.transaction DROP CONSTRAINT IF EXISTS transaction_vehicle_id_fkey;
ALTER TABLE IF EXISTS ONLY public.transaction DROP CONSTRAINT IF EXISTS transaction_employee_id_fkey;
ALTER TABLE IF EXISTS ONLY public.transaction DROP CONSTRAINT IF EXISTS transaction_customer_id_fkey;
ALTER TABLE IF EXISTS ONLY public.employee DROP CONSTRAINT IF EXISTS employee_manager_id_fkey;
ALTER TABLE IF EXISTS ONLY public.vehicle DROP CONSTRAINT IF EXISTS vehicle_pkey;
ALTER TABLE IF EXISTS ONLY public.transaction DROP CONSTRAINT IF EXISTS transaction_pkey;
ALTER TABLE IF EXISTS ONLY public.employee DROP CONSTRAINT IF EXISTS employee_pkey;
ALTER TABLE IF EXISTS ONLY public.customer DROP CONSTRAINT IF EXISTS customer_pkey;
ALTER TABLE IF EXISTS public.vehicle ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.transaction ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.employee ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.customer ALTER COLUMN id DROP DEFAULT;
DROP SEQUENCE IF EXISTS public.vehicle_id_seq;
DROP TABLE IF EXISTS public.vehicle;
DROP SEQUENCE IF EXISTS public.transaction_id_seq;
DROP TABLE IF EXISTS public.transaction;
DROP SEQUENCE IF EXISTS public.seq_vehicle;
DROP SEQUENCE IF EXISTS public.id_increment;
DROP SEQUENCE IF EXISTS public.employee_id_seq;
DROP TABLE IF EXISTS public.employee;
DROP SEQUENCE IF EXISTS public.customer_id_seq;
DROP TABLE IF EXISTS public.customer;
DROP TYPE IF EXISTS public.vehicle_status;
DROP TYPE IF EXISTS public.transaction_type;
DROP EXTENSION IF EXISTS "uuid-ossp";
--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: transaction_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.transaction_type AS ENUM (
    'Online',
    'Offline'
);


ALTER TYPE public.transaction_type OWNER TO postgres;

--
-- Name: vehicle_status; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.vehicle_status AS ENUM (
    'Baru',
    'Bekas'
);


ALTER TYPE public.vehicle_status OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: customer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customer (
    id character varying(60) NOT NULL,
    first_name character varying(30),
    last_name character varying(30),
    address character varying(225),
    phone_number character varying(20),
    email character varying(100),
    bod date
);


ALTER TABLE public.customer OWNER TO postgres;

--
-- Name: customer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customer_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customer_id_seq OWNER TO postgres;

--
-- Name: customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.customer_id_seq OWNED BY public.customer.id;


--
-- Name: employee; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.employee (
    id character varying(60) NOT NULL,
    first_name character varying(30),
    last_name character varying(30),
    address character varying(225),
    phone_number character varying(20),
    email character varying(100),
    "position" character varying(50),
    salary bigint,
    bod date,
    manager_id character varying(60)
);


ALTER TABLE public.employee OWNER TO postgres;

--
-- Name: employee_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.employee_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.employee_id_seq OWNER TO postgres;

--
-- Name: employee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.employee_id_seq OWNED BY public.employee.id;


--
-- Name: id_increment; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.id_increment
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.id_increment OWNER TO postgres;

--
-- Name: seq_vehicle; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.seq_vehicle
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.seq_vehicle OWNER TO postgres;

--
-- Name: transaction; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction (
    id character varying(60) NOT NULL,
    transaction_date date,
    vehicle_id character varying(60),
    customer_id character varying(60),
    employee_id character varying(60),
    type public.transaction_type,
    payment_amount integer,
    qty integer
);


ALTER TABLE public.transaction OWNER TO postgres;

--
-- Name: transaction_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transaction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transaction_id_seq OWNER TO postgres;

--
-- Name: transaction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transaction_id_seq OWNED BY public.transaction.id;


--
-- Name: vehicle; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vehicle (
    id character varying(60) NOT NULL,
    brand character varying(30),
    model character varying(50),
    production_year integer,
    color character varying(50),
    is_automatic boolean,
    sale_price integer NOT NULL,
    stock integer,
    status public.vehicle_status,
    CONSTRAINT sale_price_cannot_zero CHECK ((sale_price > 0))
);


ALTER TABLE public.vehicle OWNER TO postgres;

--
-- Name: vehicle_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vehicle_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vehicle_id_seq OWNER TO postgres;

--
-- Name: vehicle_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vehicle_id_seq OWNED BY public.vehicle.id;


--
-- Name: customer id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer ALTER COLUMN id SET DEFAULT nextval('public.customer_id_seq'::regclass);


--
-- Name: employee id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee ALTER COLUMN id SET DEFAULT nextval('public.employee_id_seq'::regclass);


--
-- Name: transaction id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction ALTER COLUMN id SET DEFAULT nextval('public.transaction_id_seq'::regclass);


--
-- Name: vehicle id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle ALTER COLUMN id SET DEFAULT nextval('public.vehicle_id_seq'::regclass);


--
-- Data for Name: customer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customer (id, first_name, last_name, address, phone_number, email, bod) FROM stdin;
1	Frederigo	Dabes	00 Annamark Way	5719503507	fdabesj@hao123.com	1997-09-03
2	Agatha	Fennessy	3 Crownhardt Parkway	2539210064	afennessy0@phpbb.com	1991-08-07
3	Consuelo	Hannabus	802 Anzinger Circle	6879988586	channabus1@goo.gl	1992-12-21
4	Gabriell	Kidd	69529 Lerdahl Point	5424455934	gkidd2@webeden.co.uk	1997-01-07
5	Cecilio	Blackmuir	4 Vermont Court	7775273076	cblackmuir3@printfriendly.com	1990-04-29
6	Gusty	Snary	0459 Katie Road	6626897920	gsnary4@dailymail.co.uk	1993-04-05
7	Birk	McKendo	86 Kinsman Center	7738755454	bmckendo5@slashdot.org	1999-06-18
8	Oberon	Lindermann	11895 Crownhardt Pass	9078727502	olindermann6@goo.ne.jp	1995-06-14
9	Jarrad	Klinck	56144 Cascade Place	5219999583	jklinck7@rediff.com	1997-10-27
10	Nataline	Kitchinham	38 Butternut Court	5748981969	nkitchinham8@sciencedirect.com	1995-11-03
11	Del	Mostyn	001 Evergreen Plaza	9205819890	dmostyn9@timesonline.co.uk	1992-04-09
12	Rozelle	Wanstall	5097 Fuller Park	9395944874	rwanstalla@wikia.com	1993-09-13
13	Rosemaria	Philipsen	0313 Forest Run Pass	7066808543	rphilipsenb@nps.gov	1992-06-17
14	Dani	Bloys	3 Jackson Crossing	6273622958	dbloysc@exblog.jp	1998-04-04
15	Tommy	Brambley	7900 Darwin Hill	9091350863	tbrambleyd@diigo.com	1996-04-06
16	Frederic	Ruthen	02 Laurel Trail	6702220162	fruthene@microsoft.com	1990-05-25
17	Tabbatha	MacAirt	466 Farragut Junction	6012829314	tmacairtf@mediafire.com	1992-02-06
18	Cori	Pfeffel	849 Hoard Junction	8421092375	cpfeffelg@columbia.edu	2000-02-01
19	Iain	Gladdin	3 Summer Ridge Avenue	6563272501	igladdinh@archive.org	1996-01-16
20	Filberte	Knee	15340 Spenser Place	9462597093	fkneei@yale.edu	1995-04-03
21	Rosella	Nix	0 Kim Way	4995284643	rnixj@clickbank.net	1997-01-04
afc49d21-a381-42f0-8f0b-d94d4148d8e1	Tika	Yesi		0821444444	tika.yesi@gmail.com	1999-11-11
\.


--
-- Data for Name: employee; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.employee (id, first_name, last_name, address, phone_number, email, "position", salary, bod, manager_id) FROM stdin;
17	Nicola	Burnep	5140 4th Trail	2984418544	nburnepg@vimeo.com	VP Marketing	9000000	1992-11-15	13
18	Huey	Brannon	09 Lunder Terrace	4643417218	hbrannonh@yellowpages.com	Payment Adjustment Coordinator	7500000	1998-03-14	13
19	Coreen	Dionisetti	59021 Bartelt Crossing	3742005002	cdionisettii@liveinternet.ru	VP Accounting	7500000	1993-11-16	13
21	Agatha	Fennessy	3 Crownhardt Parkway	2539210064	afennessy0@phpbb.com	\N	\N	1991-08-07	\N
22	Consuelo	Hannabus	802 Anzinger Circle	6879988586	channabus1@goo.gl	\N	\N	1992-12-21	\N
24	Del	Mostyn	001 Evergreen Plaza	9205819890	dmostyn9@timesonline.co.uk	\N	\N	1992-04-09	\N
25	Rosemaria	Philipsen	0313 Forest Run Pass	7066808543	rphilipsenb@nps.gov	\N	\N	1992-06-17	\N
27	Tabbatha	MacAirt	466 Farragut Junction	6012829314	tmacairtf@mediafire.com	\N	\N	1992-02-06	\N
23	Cecilio	Blackmuir	4 Vermont Court	7775273076	cblackmuir3@printfriendly.com	\N	\N	1990-04-29	\N
26	Frederic	Ruthen	02 Laurel Trail	6702220162	fruthene@microsoft.com	\N	\N	1990-05-25	\N
12	Deana	Pruvost	1976 Evergreen Park	8884340976	dpruvostb@mashable.com	Chief Design Engineer	25000000	1993-01-04	\N
13	Susanne	Hollyer	23 Rigney Drive	4467758188	shollyerc@cyberchimps.com	Marketing Manager	25000000	1992-03-25	\N
20	Trever	Fearnsides	4812 Banding Center	1677875665	tfearnsidesj@oakley.com	General Manager	25000000	1993-09-21	\N
3	Sebastiano	Cathel	4 Iowa Center	9561809744	scathel2@phpbb.com	Senior Editor	10000000	1996-03-22	12
4	Gladys	Crocetti	271 Mariners Cove Lane	6634704462	gcrocetti3@ifeng.com	Software Test Engineer IV	13500000	1993-03-31	20
5	Neila	Warrener	93 Mandrake Avenue	5781618978	nwarrener4@google.com.br	Financial Analyst	13000000	1996-04-18	20
6	Frasier	Chominski	0 Ramsey Court	2211575186	fchominski5@webeden.co.uk	Accounting Assistant I	9000000	1999-12-16	20
7	Nikolaos	Bottleson	39 Debra Alley	9533898988	nbottleson6@taobao.com	Senior Financial Analyst	17500000	1993-08-28	20
8	Haley	Kernermann	5971 School Drive	1567407012	hkernermann7@boston.com	Tax Accountant	8500000	1997-04-21	20
9	Boy	Queste	27 Merchant Terrace	7356247772	bqueste8@cam.ac.uk	Financial Analyst	4500000	1994-05-03	20
10	Casandra	MacKean	74001 Truax Way	6549176540	cmackean9@a8.net	Staff Accountant II	5000000	1994-04-25	20
11	Tabby	Casiroli	674 Grover Way	4847203922	tcasirolia@theguardian.com	Help Desk Operator	5000000	1994-03-31	20
14	Nicolle	Harkess	1036 Arrowood Hill	7938204869	nharkessd@live.com	Quality Control Specialist	12500000	1999-11-07	20
15	Jerry	Milson	1 South Center	3856604011	jmilsone@state.gov	Senior Quality Engineer	15000000	1993-01-06	20
16	Maryjane	Ceci	82107 Dottie Crossing	3138512748	mcecif@techcrunch.com	Marketing Manager	20000000	1993-05-18	13
2	Camellia	Gibbard	24 Moland Park	3555376689	cgibbard1@ftc.gov	Community Outreach Specialist	8500000	1991-01-15	13
1	Cletus	Menego	57 Porter Court	1353296807	cmenego0@bloomberg.com	Teacher	6500000	1996-08-29	20
34258ecc-b35c-4da9-8574-c452475af11f	Rifqi	Ramadhan		0821444244	rifqi.ramadhan@gmail.com	CTO	25000000	1990-11-11	\N
15c68c8f-eff0-42cc-a8dd-903be384fa8a	Tika	Yesi		0821444444	tika.yesi@gmail.com	Software Developer	15000000	1990-11-11	34258ecc-b35c-4da9-8574-c452475af11f
9b3f7fe7-d78b-459c-b3c5-9f220e4e04ac	Doni	Octa		082929193738	doni.octo@gmail.com	Software Developer	15000000	1990-11-11	34258ecc-b35c-4da9-8574-c452475af11f
\.


--
-- Data for Name: transaction; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transaction (id, transaction_date, vehicle_id, customer_id, employee_id, type, payment_amount, qty) FROM stdin;
T0001	0001-01-01	b3a41ff7-a5af-4f04-b0e5-19e7451a8556	afc49d21-a381-42f0-8f0b-d94d4148d8e1	15c68c8f-eff0-42cc-a8dd-903be384fa8a	Online	900000000	1
7c95d626-16a4-41fb-8ded-a54cf4e58f94	0001-01-01	b3a41ff7-a5af-4f04-b0e5-19e7451a8556	afc49d21-a381-42f0-8f0b-d94d4148d8e1	15c68c8f-eff0-42cc-a8dd-903be384fa8a	Online	900000000	1
2fb8d033-be33-4914-8b43-85e89e109778	0001-01-01	b3a41ff7-a5af-4f04-b0e5-19e7451a8556	afc49d21-a381-42f0-8f0b-d94d4148d8e1	15c68c8f-eff0-42cc-a8dd-903be384fa8a	Online	900000000	3
\.


--
-- Data for Name: vehicle; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.vehicle (id, brand, model, production_year, color, is_automatic, sale_price, stock, status) FROM stdin;
1	Honda	HR-V	2022	Putih	t	301000000	10	Baru
2	Honda	Civic	2021	Hitam	t	352000000	10	Bekas
3	Mitshubishi	XPander	2020	Silver Metalik	f	231750000	10	Baru
4	Toyota	Rush	2021	Hitam	f	232000000	10	Bekas
5	Mazda	CX-3	2022	Putih	t	302000000	10	Baru
6	Toyota	Sienta	2020	Merah	t	198500000	10	Bekas
7	Toyota	Yaris	2020	Putih	f	171000000	10	Bekas
8	BMX	X1	2022	Hitam	f	443000000	10	Baru
9	Nissan	Serena	2019	Putih	t	177000000	10	Baru
10	Mitshubishi	Pajero Sport	2022	Putih	t	470000000	10	Bekas
11	Mazda	CX-9	2022	Hitam	t	552000000	10	Baru
12	BMW	3i	2022	Putih	f	480000000	10	Baru
13	Datsun	Go-Panca T 1.2	2015	Putih	f	77000000	10	Bekas
14	Suzuki	Ertiga	2019	Hitam	t	196000000	10	Baru
15	Suzuki	Ertiga GX 1.5	2022	Hitam	t	227000000	10	Baru
16	Mercedes-Benz	GLA 200 AMG 1.6	2018	Hitam	f	569000000	10	Bekas
17	Mercedes-Benz	C 300 Coupe AMG 2.0	2016	Hitam	f	828000000	10	Bekas
b3a41ff7-a5af-4f04-b0e5-19e7451a8556	Toyota	Alphard	0	Putih	f	900000000	0	Baru
c7ea13da-8fdb-4275-b86a-97f7b5a8aa24	Toyota	Alphard	0	Putih	f	900000000	0	Baru
\.


--
-- Name: customer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.customer_id_seq', 21, true);


--
-- Name: employee_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.employee_id_seq', 27, true);


--
-- Name: id_increment; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.id_increment', 1, false);


--
-- Name: seq_vehicle; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.seq_vehicle', 25, true);


--
-- Name: transaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transaction_id_seq', 30, true);


--
-- Name: vehicle_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vehicle_id_seq', 17, true);


--
-- Name: customer customer_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer
    ADD CONSTRAINT customer_pkey PRIMARY KEY (id);


--
-- Name: employee employee_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_pkey PRIMARY KEY (id);


--
-- Name: transaction transaction_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);


--
-- Name: vehicle vehicle_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle
    ADD CONSTRAINT vehicle_pkey PRIMARY KEY (id);


--
-- Name: employee employee_manager_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.employee
    ADD CONSTRAINT employee_manager_id_fkey FOREIGN KEY (manager_id) REFERENCES public.employee(id);


--
-- Name: transaction transaction_customer_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_customer_id_fkey FOREIGN KEY (customer_id) REFERENCES public.customer(id);


--
-- Name: transaction transaction_employee_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES public.employee(id);


--
-- Name: transaction transaction_vehicle_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction
    ADD CONSTRAINT transaction_vehicle_id_fkey FOREIGN KEY (vehicle_id) REFERENCES public.vehicle(id);


--
-- PostgreSQL database dump complete
--

