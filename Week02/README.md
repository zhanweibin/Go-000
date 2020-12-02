Week02 作业题目：

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：dao层属于应用业务层，对于底层kit库报的error, 应该wrap这个error抛给上层，并在最上层处理。 