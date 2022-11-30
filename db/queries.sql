-- name: GetLinuxUsers :many
SELECT * FROM public.linux_user 
order by linux_user_name;

-- name: GetLinuxUser :one
SELECT * from public.linux_user 
 where linux_user_id = $1
LIMIT 1; 

-- name: CreateLinuxUser :one
INSERT INTO public.linux_user (
linux_user_name, linux_distro
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteLinuxUser :exec
DELETE FROM public.linux_user 
WHERE linux_user_id = $1;

-- name: UpdateLinuxUser :exec
UPDATE public.linux_user 
  set 
      linux_user_name = $2, 
      linux_distro = $3
WHERE linux_user_id  = $1;

-- name: GetLinuxDistros :many
SELECT linux_distro FROM public.linux_distro
order by linux_distro;


-- name: GetLinuxDistro :one
SELECT linux_distro FROM public.linux_distro
where linux_distro = $1
limit 1;
