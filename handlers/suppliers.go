package handlers

//const (
//	TopicSuppliersUpserted = "suppliers:upserted"
//)
//
//type Suppliers struct {
//	db       db.Suppliers
//	stanConn stan.Conn
//}
//
//func NewSuppliersHandler(db db.Suppliers, stanConn stan.Conn) *Suppliers {
//	return &Suppliers{db: db, stanConn: stanConn}
//}
//
//func (h *Suppliers) UpsertSuppliers(ctx context.Context, req *psuppliers.UpsertSuppliersReq, rsp *psuppliers.UpsertSuppliersRsp) error {
//	// bulk upserts
//	if err := h.db.UpsertSuppliers(ctx, req.Suppliers); err != nil {
//		return err
//	}
//
//	// get all suppliers and construct response
//	suppliers, err := h.db.GetAllSuppliers(ctx)
//	if err != nil {
//		return err
//	}
//	rsp.Suppliers = suppliers
//
//	// publish message
//	rspBytes, err := proto.Marshal(rsp)
//	if err != nil {
//		return err
//	}
//	if err := h.stanConn.Publish(TopicSuppliersUpserted, rspBytes); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (h *Suppliers) Suppliers(ctx context.Context, req *psuppliers.SuppliersReq, rsp *psuppliers.SuppliersRsp) error {
//	// get all suppliers and construct response
//	suppliers, err := h.db.GetAllSuppliers(ctx)
//	if err != nil {
//		return err
//	}
//	rsp.Suppliers = suppliers
//
//	return nil
//}
