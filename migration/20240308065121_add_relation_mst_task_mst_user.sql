-- +goose Up
-- +goose StatementBegin
ALTER TABLE mst_task
ADD CONSTRAINT `FK_mst_task_mst_user` 
FOREIGN KEY (`user_id`) REFERENCES `mst_user` (`id`);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE mst_task DROP FOREIGN KEY FK_mst_task_mst_user;
-- +goose StatementEnd
