PGDMP         6            	    w            golang    9.2.24    9.2.24 E    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                       false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                       false            �           1262    16384    golang    DATABASE     x   CREATE DATABASE golang WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
    DROP DATABASE golang;
             postgres    false                        2615    2200    public    SCHEMA        CREATE SCHEMA public;
    DROP SCHEMA public;
             postgres    false            �           0    0    SCHEMA public    COMMENT     6   COMMENT ON SCHEMA public IS 'standard public schema';
                  postgres    false    6            �           0    0    public    ACL     �   REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;
                  postgres    false    6                        3079    12649    plpgsql 	   EXTENSION     ?   CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
    DROP EXTENSION plpgsql;
                  false            �           0    0    EXTENSION plpgsql    COMMENT     @   COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
                       false    1            �            1259    16597    admin    TABLE     �  CREATE TABLE admin (
    id integer NOT NULL,
    email character varying(255),
    name character varying(255),
    password character varying(255),
    role_id bigint,
    last_login_ip character varying(32),
    last_login_time character varying(255),
    login_count integer,
    is_lock boolean DEFAULT false,
    describe character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
    DROP TABLE public.admin;
       public         postgres    false    6            �            1259    16595    admin_id_seq    SEQUENCE     n   CREATE SEQUENCE admin_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.admin_id_seq;
       public       postgres    false    6    174            �           0    0    admin_id_seq    SEQUENCE OWNED BY     /   ALTER SEQUENCE admin_id_seq OWNED BY admin.id;
            public       postgres    false    173            �            1259    16572    admin_navigation    TABLE     A  CREATE TABLE admin_navigation (
    id integer NOT NULL,
    url character varying(255),
    title character varying(255),
    parent_id integer,
    is_show boolean DEFAULT true,
    is_sys boolean DEFAULT true,
    sort integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
 $   DROP TABLE public.admin_navigation;
       public         postgres    false    6            �            1259    16570    admin_navigation_id_seq    SEQUENCE     y   CREATE SEQUENCE admin_navigation_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.admin_navigation_id_seq;
       public       postgres    false    170    6            �           0    0    admin_navigation_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE admin_navigation_id_seq OWNED BY admin_navigation.id;
            public       postgres    false    169            �            1259    16586    admin_navigation_node    TABLE       CREATE TABLE admin_navigation_node (
    id integer NOT NULL,
    admin_navigation_id integer,
    route_action character varying(255),
    title character varying(255),
    sort integer,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
 )   DROP TABLE public.admin_navigation_node;
       public         postgres    false    6            �            1259    16584    admin_navigation_node_id_seq    SEQUENCE     ~   CREATE SEQUENCE admin_navigation_node_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 3   DROP SEQUENCE public.admin_navigation_node_id_seq;
       public       postgres    false    172    6            �           0    0    admin_navigation_node_id_seq    SEQUENCE OWNED BY     O   ALTER SEQUENCE admin_navigation_node_id_seq OWNED BY admin_navigation_node.id;
            public       postgres    false    171            �            1259    16647    article    TABLE     �  CREATE TABLE article (
    id integer NOT NULL,
    title character varying(255),
    is_show boolean DEFAULT true,
    sort integer NOT NULL,
    abstract character varying(255),
    author character varying(255),
    content text,
    cover character varying(255),
    category_id integer DEFAULT 0,
    tags character varying(255),
    sources character varying(255),
    allow_comments boolean DEFAULT true,
    top smallint,
    status smallint DEFAULT 1,
    view_num integer DEFAULT 0,
    comment_num integer DEFAULT 0,
    likes_num integer DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
    DROP TABLE public.article;
       public         postgres    false    6            �            1259    16645    article_id_seq    SEQUENCE     p   CREATE SEQUENCE article_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.article_id_seq;
       public       postgres    false    6    182            �           0    0    article_id_seq    SEQUENCE OWNED BY     3   ALTER SEQUENCE article_id_seq OWNED BY article.id;
            public       postgres    false    181            �            1259    16633    category    TABLE     *  CREATE TABLE category (
    id integer NOT NULL,
    title character varying(255),
    is_show boolean DEFAULT true,
    sort integer NOT NULL,
    describe character varying(255),
    pid integer DEFAULT 0,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
    DROP TABLE public.category;
       public         postgres    false    6            �            1259    16631    category_id_seq    SEQUENCE     q   CREATE SEQUENCE category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public.category_id_seq;
       public       postgres    false    180    6            �           0    0    category_id_seq    SEQUENCE OWNED BY     5   ALTER SEQUENCE category_id_seq OWNED BY category.id;
            public       postgres    false    179            �            1259    16611    role    TABLE       CREATE TABLE role (
    id integer NOT NULL,
    role_name character varying(255),
    is_super boolean DEFAULT false,
    is_sys boolean DEFAULT true,
    describe character varying(255),
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
    DROP TABLE public.role;
       public         postgres    false    6            �            1259    16609    role_id_seq    SEQUENCE     m   CREATE SEQUENCE role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.role_id_seq;
       public       postgres    false    176    6            �           0    0    role_id_seq    SEQUENCE OWNED BY     -   ALTER SEQUENCE role_id_seq OWNED BY role.id;
            public       postgres    false    175            �            1259    16625 	   role_node    TABLE     �   CREATE TABLE role_node (
    id integer NOT NULL,
    role_id integer NOT NULL,
    admin_navigation_id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
    DROP TABLE public.role_node;
       public         postgres    false    6            �            1259    16623    role_node_id_seq    SEQUENCE     r   CREATE SEQUENCE role_node_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.role_node_id_seq;
       public       postgres    false    178    6            �           0    0    role_node_id_seq    SEQUENCE OWNED BY     7   ALTER SEQUENCE role_node_id_seq OWNED BY role_node.id;
            public       postgres    false    177            �            1259    16665    role_node_routes    TABLE     �   CREATE TABLE role_node_routes (
    id integer NOT NULL,
    role_id integer NOT NULL,
    admin_navigation_node_id integer NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone
);
 $   DROP TABLE public.role_node_routes;
       public         postgres    false    6            �            1259    16663    role_node_routes_id_seq    SEQUENCE     y   CREATE SEQUENCE role_node_routes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.role_node_routes_id_seq;
       public       postgres    false    184    6            �           0    0    role_node_routes_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE role_node_routes_id_seq OWNED BY role_node_routes.id;
            public       postgres    false    183            �
           2604    16600    id    DEFAULT     V   ALTER TABLE ONLY admin ALTER COLUMN id SET DEFAULT nextval('admin_id_seq'::regclass);
 7   ALTER TABLE public.admin ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    173    174    174            �
           2604    16575    id    DEFAULT     l   ALTER TABLE ONLY admin_navigation ALTER COLUMN id SET DEFAULT nextval('admin_navigation_id_seq'::regclass);
 B   ALTER TABLE public.admin_navigation ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    170    169    170            �
           2604    16589    id    DEFAULT     v   ALTER TABLE ONLY admin_navigation_node ALTER COLUMN id SET DEFAULT nextval('admin_navigation_node_id_seq'::regclass);
 G   ALTER TABLE public.admin_navigation_node ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    172    171    172            �
           2604    16650    id    DEFAULT     Z   ALTER TABLE ONLY article ALTER COLUMN id SET DEFAULT nextval('article_id_seq'::regclass);
 9   ALTER TABLE public.article ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    181    182    182            �
           2604    16636    id    DEFAULT     \   ALTER TABLE ONLY category ALTER COLUMN id SET DEFAULT nextval('category_id_seq'::regclass);
 :   ALTER TABLE public.category ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    180    179    180            �
           2604    16614    id    DEFAULT     T   ALTER TABLE ONLY role ALTER COLUMN id SET DEFAULT nextval('role_id_seq'::regclass);
 6   ALTER TABLE public.role ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    175    176    176            �
           2604    16628    id    DEFAULT     ^   ALTER TABLE ONLY role_node ALTER COLUMN id SET DEFAULT nextval('role_node_id_seq'::regclass);
 ;   ALTER TABLE public.role_node ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    177    178    178            �
           2604    16668    id    DEFAULT     l   ALTER TABLE ONLY role_node_routes ALTER COLUMN id SET DEFAULT nextval('role_node_routes_id_seq'::regclass);
 B   ALTER TABLE public.role_node_routes ALTER COLUMN id DROP DEFAULT;
       public       postgres    false    183    184    184            t          0    16597    admin 
   TABLE DATA               �   COPY admin (id, email, name, password, role_id, last_login_ip, last_login_time, login_count, is_lock, describe, created_at, updated_at) FROM stdin;
    public       postgres    false    174   tL       �           0    0    admin_id_seq    SEQUENCE SET     3   SELECT pg_catalog.setval('admin_id_seq', 2, true);
            public       postgres    false    173            p          0    16572    admin_navigation 
   TABLE DATA               m   COPY admin_navigation (id, url, title, parent_id, is_show, is_sys, sort, created_at, updated_at) FROM stdin;
    public       postgres    false    170   M       �           0    0    admin_navigation_id_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('admin_navigation_id_seq', 1, true);
            public       postgres    false    169            r          0    16586    admin_navigation_node 
   TABLE DATA               t   COPY admin_navigation_node (id, admin_navigation_id, route_action, title, sort, created_at, updated_at) FROM stdin;
    public       postgres    false    172   BM       �           0    0    admin_navigation_node_id_seq    SEQUENCE SET     C   SELECT pg_catalog.setval('admin_navigation_node_id_seq', 1, true);
            public       postgres    false    171            |          0    16647    article 
   TABLE DATA               �   COPY article (id, title, is_show, sort, abstract, author, content, cover, category_id, tags, sources, allow_comments, top, status, view_num, comment_num, likes_num, created_at, updated_at) FROM stdin;
    public       postgres    false    182   �M       �           0    0    article_id_seq    SEQUENCE SET     6   SELECT pg_catalog.setval('article_id_seq', 1, false);
            public       postgres    false    181            z          0    16633    category 
   TABLE DATA               \   COPY category (id, title, is_show, sort, describe, pid, created_at, updated_at) FROM stdin;
    public       postgres    false    180   �M       �           0    0    category_id_seq    SEQUENCE SET     6   SELECT pg_catalog.setval('category_id_seq', 2, true);
            public       postgres    false    179            v          0    16611    role 
   TABLE DATA               Z   COPY role (id, role_name, is_super, is_sys, describe, created_at, updated_at) FROM stdin;
    public       postgres    false    176    N       �           0    0    role_id_seq    SEQUENCE SET     2   SELECT pg_catalog.setval('role_id_seq', 2, true);
            public       postgres    false    175            x          0    16625 	   role_node 
   TABLE DATA               V   COPY role_node (id, role_id, admin_navigation_id, created_at, updated_at) FROM stdin;
    public       postgres    false    178   xN       �           0    0    role_node_id_seq    SEQUENCE SET     8   SELECT pg_catalog.setval('role_node_id_seq', 1, false);
            public       postgres    false    177            ~          0    16665    role_node_routes 
   TABLE DATA               b   COPY role_node_routes (id, role_id, admin_navigation_node_id, created_at, updated_at) FROM stdin;
    public       postgres    false    184   �N       �           0    0    role_node_routes_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('role_node_routes_id_seq', 1, false);
            public       postgres    false    183            �
           2606    16594    admin_navigation_node_pkey 
   CONSTRAINT     g   ALTER TABLE ONLY admin_navigation_node
    ADD CONSTRAINT admin_navigation_node_pkey PRIMARY KEY (id);
 Z   ALTER TABLE ONLY public.admin_navigation_node DROP CONSTRAINT admin_navigation_node_pkey;
       public         postgres    false    172    172            �
           2606    16582    admin_navigation_pkey 
   CONSTRAINT     ]   ALTER TABLE ONLY admin_navigation
    ADD CONSTRAINT admin_navigation_pkey PRIMARY KEY (id);
 P   ALTER TABLE ONLY public.admin_navigation DROP CONSTRAINT admin_navigation_pkey;
       public         postgres    false    170    170            �
           2606    16606 
   admin_pkey 
   CONSTRAINT     G   ALTER TABLE ONLY admin
    ADD CONSTRAINT admin_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.admin DROP CONSTRAINT admin_pkey;
       public         postgres    false    174    174                       2606    16662    article_pkey 
   CONSTRAINT     K   ALTER TABLE ONLY article
    ADD CONSTRAINT article_pkey PRIMARY KEY (id);
 >   ALTER TABLE ONLY public.article DROP CONSTRAINT article_pkey;
       public         postgres    false    182    182                        2606    16643    category_pkey 
   CONSTRAINT     M   ALTER TABLE ONLY category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);
 @   ALTER TABLE ONLY public.category DROP CONSTRAINT category_pkey;
       public         postgres    false    180    180            �
           2606    16630    role_node_pkey 
   CONSTRAINT     O   ALTER TABLE ONLY role_node
    ADD CONSTRAINT role_node_pkey PRIMARY KEY (id);
 B   ALTER TABLE ONLY public.role_node DROP CONSTRAINT role_node_pkey;
       public         postgres    false    178    178                       2606    16670    role_node_routes_pkey 
   CONSTRAINT     ]   ALTER TABLE ONLY role_node_routes
    ADD CONSTRAINT role_node_routes_pkey PRIMARY KEY (id);
 P   ALTER TABLE ONLY public.role_node_routes DROP CONSTRAINT role_node_routes_pkey;
       public         postgres    false    184    184            �
           2606    16621 	   role_pkey 
   CONSTRAINT     E   ALTER TABLE ONLY role
    ADD CONSTRAINT role_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.role DROP CONSTRAINT role_pkey;
       public         postgres    false    176    176            �
           1259    16583    IDX_admin_navigation_title    INDEX     S   CREATE INDEX "IDX_admin_navigation_title" ON admin_navigation USING btree (title);
 0   DROP INDEX public."IDX_admin_navigation_title";
       public         postgres    false    170            �
           1259    16608    IDX_admin_role_id    INDEX     A   CREATE INDEX "IDX_admin_role_id" ON admin USING btree (role_id);
 '   DROP INDEX public."IDX_admin_role_id";
       public         postgres    false    174            �
           1259    16607    UQE_admin_name    INDEX     B   CREATE UNIQUE INDEX "UQE_admin_name" ON admin USING btree (name);
 $   DROP INDEX public."UQE_admin_name";
       public         postgres    false    174            �
           1259    16644    UQE_category_title    INDEX     J   CREATE UNIQUE INDEX "UQE_category_title" ON category USING btree (title);
 (   DROP INDEX public."UQE_category_title";
       public         postgres    false    180            �
           1259    16622    UQE_role_role_name    INDEX     J   CREATE UNIQUE INDEX "UQE_role_role_name" ON role USING btree (role_name);
 (   DROP INDEX public."UQE_role_role_name";
       public         postgres    false    176            t   }   x�����0C��W���@r����ԡE��B������7ܣ|��uk�u���^G�����ew�)KM��E��P�t�
�
Nȕ�MT�6� F:��	E�Oae"��Fu����4�<g      p   1   x�3��444�4�LBKKN#CK]C]C#Cs++slb\1z\\\ E"
�      r   V   x�3�4�L,(�O��+)���I-*�OJL�N/�/�K�sO-	-N-��,.�|�1���@�F������F
��V&&V��ĸb���� 4:      |      x������ � �      z   K   x�3��I,J,K��,�4���Mk�z��������������������)61.#��|�	 
�f#clb\1z\\\ 5�"5      v   H   x�3�|��������-|>�����%@�id`h�k`�kh�`hnedfel�M�ˈ���u/f!����=F��� ��'S      x      x������ � �      ~      x������ � �     